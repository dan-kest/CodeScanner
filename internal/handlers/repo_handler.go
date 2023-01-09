package handlers

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/handlers/payloads"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

type RepoHandler struct {
	conf        *config.Config
	qConn       *amqp091.Connection
	repoService *services.RepoService
}

func NewRepoHandler(conf *config.Config, qConn *amqp091.Connection, repoService *services.RepoService) *RepoHandler {
	return &RepoHandler{
		conf:        conf,
		qConn:       qConn,
		repoService: repoService,
	}
}

func (h *RepoHandler) ListRepo(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "'page' must be a number")
	}

	itemStr := ctx.Query("item")
	item, err := strconv.Atoi(itemStr)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "'page' must be a number")
	}

	paging := &models.Paging{
		Page:        page,
		ItemPerPage: item,
	}

	repoPagination, err := h.repoService.ListRepo(paging)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	itemList := []*payloads.ListRepoResponseItem{}
	for _, item := range repoPagination.ItemList {
		var scanStatus *string
		var timestamp *string

		if item.ScanStatus != "" {
			scanStatus = (*string)(&item.ScanStatus)
		}
		if item.Timestamp != nil {
			s := item.Timestamp.Format(time.RFC3339)
			timestamp = &s
		}
		itemList = append(itemList, &payloads.ListRepoResponseItem{
			ID:         item.ID.String(),
			Name:       *item.Name,
			URL:        *item.URL,
			ScanStatus: scanStatus,
			Timestamp:  timestamp,
		})
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
		Data: payloads.ListRepoResponse{
			Page:        repoPagination.Page,
			ItemPerPage: repoPagination.ItemPerPage,
			TotalCount:  repoPagination.TotalCount,
			ItemList:    itemList,
		},
	})
}

func (h *RepoHandler) ViewRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "invalid id")
	}

	repo, err := h.repoService.ViewRepo(id)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if strings.HasSuffix(err.Error(), constants.ErrorNotFoundSuffix) {
			statusCode = fiber.StatusNotFound
		}

		return sendError(ctx, statusCode, err.Error())
	}

	var findingList []*payloads.Finding
	if len(repo.Findings) > 0 {
		findingList = []*payloads.Finding{}
		for _, finding := range repo.Findings {
			findingList = append(findingList, &payloads.Finding{
				Type:   finding.Type,
				RuleID: finding.RuleID,
				Location: payloads.FindingLocation{
					Path: finding.Location.Path,
					Positions: payloads.FindingLocationPosition{
						Begin: payloads.FindingLocationPositionBegin{
							Line: finding.Location.Positions.Begin.Line,
						},
					},
				},
				Metadata: payloads.FindingMetadata{
					Description: finding.Metadata.Description,
					Severity:    finding.Metadata.Severity,
				},
			})
		}
	}
	var scanStatus *string
	var timestamp *string

	if repo.ScanStatus != "" {
		scanStatus = (*string)(&repo.ScanStatus)
	}
	if repo.Timestamp != nil {
		t := repo.Timestamp.Format(time.RFC3339)
		timestamp = &t
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
		Data: payloads.ViewRepoResponse{
			ID:         repo.ID.String(),
			Name:       *repo.Name,
			URL:        *repo.URL,
			ScanStatus: scanStatus,
			Timestamp:  timestamp,
			Findings:   findingList,
		},
	})
}

func (h *RepoHandler) ScanRepo(ctx *fiber.Ctx) error {
	payload := &payloads.ScanRequest{}
	if err := ctx.BodyParser(payload); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}

	messageList := []string{}
	if payload.ID == "" {
		messageList = append(messageList, "id is required")
	}
	if payload.URL == "" {
		messageList = append(messageList, "url is required")
	}
	if len(messageList) > 0 {
		return sendError(ctx, fiber.StatusBadRequest, strings.Join(messageList, ","))
	}

	id, err := uuid.Parse(payload.ID)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "invalid id")
	}

	scanID := uuid.New()

	task := &models.Task{
		RepositoryIDStr: payload.ID,
		ScanIDStr:       scanID.String(),
		URL:             payload.URL,
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
	}
	body, err := json.Marshal(task)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}
	timeout := h.conf.RabbitMQ.PublishTimeout
	if err := publishMessage(h.qConn, h.conf.RabbitMQ.Queue.Name, body, timeout); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	if err := h.repoService.ScanRepo(id, scanID); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
	})
}

func (h *RepoHandler) CreateRepo(ctx *fiber.Ctx) error {
	payload := &payloads.RepoRequest{}
	if err := ctx.BodyParser(payload); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}

	messageList := []string{}
	if payload.Name == nil || *payload.Name == "" {
		messageList = append(messageList, "name is required")
	}
	if payload.URL == nil || *payload.URL == "" {
		messageList = append(messageList, "url is required")
	}
	if len(messageList) > 0 {
		return sendError(ctx, fiber.StatusBadRequest, strings.Join(messageList, ","))
	}

	repo := &models.Repo{
		Name: payload.Name,
		URL:  payload.URL,
	}

	id, err := h.repoService.CreateRepo(repo)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
		Data:   id.String(),
	})
}

func (h *RepoHandler) UpdateRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "invalid id")
	}

	payload := &payloads.RepoRequest{}
	if err := ctx.BodyParser(payload); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}

	repo := &models.Repo{
		Name: payload.Name,
		URL:  payload.URL,
	}

	if err := h.repoService.UpdateRepo(id, repo); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
	})
}

func (h *RepoHandler) DeleteRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.repoService.DeleteRepo(id); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(payloads.GenericResponse{
		Status: constants.ResponseStatusOK,
	})
}
