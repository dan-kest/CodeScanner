package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RepoHandler struct {
	repoService *services.RepoService
}

func NewRepoHandler(repoService *services.RepoService) *RepoHandler {
	return &RepoHandler{
		repoService: repoService,
	}
}

type RepoPayload struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

func (h *RepoHandler) ListRepo(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'page' must be a number",
		})
	}

	itemStr := ctx.Query("item")
	item, err := strconv.Atoi(itemStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'item' must be a number",
		})
	}

	paging := &models.Paging{
		Page:        page,
		ItemPerPage: item,
	}

	repoPagination, err := h.repoService.ListRepo(paging)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	repoNameListStr := ""
	comma := ""
	for _, repo := range repoPagination.ItemList {
		repoNameListStr += fmt.Sprintf("%s'%s'", comma, repo.Name.Val)
		comma = ","
	}

	return ctx.JSON(fmt.Sprintf("list: count=%d items=%s", repoPagination.TotalCount, repoNameListStr))
}

func (h *RepoHandler) ViewRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "invalid id",
		})
	}

	repo, err := h.repoService.ViewRepo(id)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if strings.HasSuffix(err.Error(), constants.ErrorNotFoundSuffix) {
			statusCode = fiber.StatusNotFound
		}

		return ctx.Status(statusCode).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON("view: " + repo.Name.Val + "; " + repo.URL.Val)
}

func (h *RepoHandler) ScanRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "invalid id",
		})
	}

	h.repoService.ScanRepo(id)

	return ctx.JSON("scan")
}

func (h *RepoHandler) CreateRepo(ctx *fiber.Ctx) error {
	payload := &RepoPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	repo := &models.Repo{}
	if payload.Name != nil {
		repo.Name.Set(*payload.Name)
	}
	if payload.URL != nil {
		repo.URL.Set(*payload.URL)
	}

	id, err := h.repoService.CreateRepo(repo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON("create: " + id.String())
}

func (h *RepoHandler) UpdateRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "invalid id",
		})
	}

	payload := &RepoPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	repo := &models.Repo{}
	if payload.Name != nil {
		repo.Name.Set(*payload.Name)
	}
	if payload.URL != nil {
		repo.URL.Set(*payload.URL)
	}

	if err := h.repoService.UpdateRepo(id, repo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON("update: OK")
}

func (h *RepoHandler) DeleteRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "invalid id",
		})
	}

	if err := h.repoService.DeleteRepo(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON("delete: OK")
}
