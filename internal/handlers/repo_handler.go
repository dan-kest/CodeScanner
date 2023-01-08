package handlers

import (
	"strconv"

	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/gofiber/fiber/v2"
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
			Message: "'page' must be a number.",
		})
	}

	itemStr := ctx.Query("item")
	item, err := strconv.Atoi(itemStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'item' must be a number.",
		})
	}

	paging := &models.Paging{
		Page:        page,
		ItemPerPage: item,
	}

	h.repoService.ListRepo(paging)

	return ctx.JSON("list")
}

func (h *RepoHandler) ViewRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'id' must be a number.",
		})
	}

	h.repoService.ViewRepo(id)

	return ctx.JSON("view")
}

func (h *RepoHandler) ScanRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'id' must be a number.",
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

	h.repoService.CreateRepo(repo)

	return ctx.JSON("create")
}

func (h *RepoHandler) UpdateRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'id' must be a number.",
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

	h.repoService.UpdateRepo(id, repo)

	return ctx.JSON("update")
}

func (h *RepoHandler) DeleteRepo(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "'id' must be a number.",
		})
	}

	h.repoService.DeleteRepo(id)

	return ctx.JSON("delete")
}
