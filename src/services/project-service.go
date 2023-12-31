package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	member "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/member-service/v1"
	project "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/project-service/v1"
	token "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/token-service/v1"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type projectService struct {
	projectClient gateways.ProjectClient
	memberClient  gateways.MemberClient
	imageClient   gateways.ImageClient
	tokenClient   gateways.TokenClient
}

type ProjectService interface {
	CreateProject(ctx context.Context, title string) (*model.Project, error)
	GetProject(ctx context.Context, projectID string) (*model.Project, error)
	GetProjects(ctx context.Context, userID string) ([]*model.Project, error)
	UpdateProject(ctx context.Context, projectID, title string, image *graphql.Upload) (*model.Project, error)
	DeleteProject(ctx context.Context, projectId string) (*string, error)
	CreateInviteLink(ctx context.Context, projectID string, authority model.Auth) (*string, error)
}

func NewProjectService(projectClient gateways.ProjectClient, memberClient gateways.MemberClient, tokenClient gateways.TokenClient, imageClient gateways.ImageClient) ProjectService {
	return &projectService{
		projectClient: projectClient,
		memberClient:  memberClient,
		tokenClient:   tokenClient,
		imageClient:   imageClient,
	}
}

func (p *projectService) CreateProject(ctx context.Context, title string) (*model.Project, error) {
	if len(title) == 0 {
		title = "untitled"
	}

	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	result, err := p.projectClient.CreateProject(ctx, &project.CreateProjectRequest{
		Title:  title,
		UserId: payload.UserId,
	})

	if err != nil {
		return nil, err
	}

	_, err = p.memberClient.CreateProjectMember(ctx, &member.MemberRequest{
		ProjectId: result.ProjectId,
		UserId:    payload.UserId,
		Authority: member.Auth.Enum(member.Auth_owner).String(),
	})

	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) GetProject(ctx context.Context, projectID string) (*model.Project, error) {
	result, err := p.projectClient.GetOneProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) GetProjects(ctx context.Context, userId string) ([]*model.Project, error) {
	result, err := p.projectClient.GetProjects(ctx, userId)
	if err != nil {
		return nil, err
	}

	projects := make([]*model.Project, 0, len(result))
	for _, project := range result {
		projects = append(projects, &model.Project{
			ProjectID: project.ProjectId,
			Title:     project.Title,
			LastImage: project.LastImage,
			UpdatedAt: time.Now().String(),
		})
	}

	return projects, nil
}

func (p *projectService) UpdateProject(ctx context.Context, projectID, title string, image *graphql.Upload) (*model.Project, error) {
	var (
		result *project.ProjectDetails
		err    error
	)

	switch {
	case len(title) != 0 && image == nil:
		result, err = p.projectClient.UpdateProjectTitle(ctx, &project.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
	case len(title) == 0 && image != nil:
		result, err = p.imageUpload(ctx, projectID, image)
	case len(title) != 0 && image != nil:
		_, err = p.projectClient.UpdateProjectTitle(ctx, &project.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
		if err != nil {
			return nil, err
		}

		result, err = p.imageUpload(ctx, projectID, image)
	default:
		return nil, fmt.Errorf("title or image must be specified")
	}

	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) imageUpload(ctx context.Context, projectID string, image *graphql.Upload) (*project.ProjectDetails, error) {
	response, err := p.imageClient.UploadImage(ctx, projectID, image)
	if err != nil {
		return nil, err
	}

	return p.projectClient.UpdateProjectImage(ctx, projectID, response.Path)
}

func (p *projectService) DeleteProject(ctx context.Context, projectId string) (*string, error) {
	project, err := p.projectClient.DeleteProject(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *projectService) CreateInviteLink(ctx context.Context, projectID string, authority model.Auth) (*string, error) {
	result, err := p.tokenClient.CreateToken(ctx, &token.CreateTokenRequest{
		ProjectId: projectID,
		Authority: token.Auth(token.Auth_value[strings.ToLower(authority.String())]),
	})
	if err != nil {
		return nil, err
	}
	return &result.Token, nil
}
