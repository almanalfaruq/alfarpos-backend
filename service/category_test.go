package service

import (
	"fmt"
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCategoryService_GetOneCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	catRepo := NewMockcategoryRepositoryIface(ctrl)
	category := CategoryService{
		category: catRepo,
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		mock    func(args args)
		args    args
		want    model.Category
		wantErr bool
	}{
		{
			name: "Success",
			mock: func(args args) {
				catRepo.EXPECT().FindById(args.id).Return(model.Category{
					Template: model.Template{
						ID: uint(args.id),
					},
					Name: "bayi",
				}, nil)
			},
			args: args{
				id: int64(20),
			},
			want: model.Category{
				Template: model.Template{
					ID: uint(20),
				},
				Name: "bayi",
			},
		},
		{
			name: "Error-NotFound",
			mock: func(args args) {
				catRepo.EXPECT().FindById(args.id).Return(model.Category{
					Template: model.Template{
						ID: uint(0),
					},
					Name: "bayi",
				}, nil)
			},
			args: args{
				id: int64(20),
			},
			want:    model.Category{},
			wantErr: true,
		},
		{
			name: "Error-DB",
			mock: func(args args) {
				catRepo.EXPECT().FindById(args.id).Return(model.Category{}, fmt.Errorf("error"))
			},
			args: args{
				id: int64(20),
			},
			want:    model.Category{},
			wantErr: true,
		},
		{
			name: "Error-Param",
			mock: func(args args) {
			},
			args: args{
				id: int64(0),
			},
			want:    model.Category{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			got, err := category.GetOneCategory(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetOneCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCategoryService_GetCategoriesByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	catRepo := NewMockcategoryRepositoryIface(ctrl)
	category := CategoryService{
		category: catRepo,
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		mock    func(name string)
		args    args
		want    []model.Category
		wantErr bool
	}{
		{
			name: "Success",
			mock: func(name string) {
				catRepo.EXPECT().FindByName(name).Return([]model.Category{
					{

						Template: model.Template{
							ID: uint(20),
						},
						Name: "perkantoran",
					},
				})
			},
			args: args{
				name: "kantor",
			},
			want: []model.Category{
				{

					Template: model.Template{
						ID: uint(20),
					},
					Name: "perkantoran",
				},
			},
		},
		{
			name: "Error",
			mock: func(name string) {
				catRepo.EXPECT().FindByName(name).Return([]model.Category{})
			},
			args: args{
				name: "kantor",
			},
			want:    []model.Category{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args.name)
			got, err := category.GetCategoriesByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetCategoriesByName() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCategoryService_NewCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	catRepo := NewMockcategoryRepositoryIface(ctrl)
	category := CategoryService{
		category: catRepo,
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		mock    func(args args)
		args    args
		want    model.Category
		wantErr bool
	}{
		{
			name: "Success",
			mock: func(args args) {
				catRepo.EXPECT().New(model.Category{
					Name: args.name,
				}).Return(model.Category{
					Template: model.Template{
						ID: uint(10),
					},
					Name: args.name,
				}, nil)
			},
			args: args{
				name: "alat tulis",
			},
			want: model.Category{
				Template: model.Template{
					ID: uint(10),
				},
				Name: "alat tulis",
			},
		},
		{
			name: "Error-Param",
			mock: func(args args) {
			},
			args: args{
				name: "",
			},
			want:    model.Category{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			got, err := category.NewCategory(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.NewCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCategoryService_DeleteCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	catRepo := NewMockcategoryRepositoryIface(ctrl)
	category := CategoryService{
		category: catRepo,
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		mock    func(args args)
		args    args
		want    model.Category
		wantErr bool
	}{
		{
			name: "Success",
			mock: func(args args) {
				catRepo.EXPECT().Delete(args.id).Return(model.Category{
					Template: model.Template{
						ID: uint(args.id),
					},
					Name: "alat tulis",
				}, nil)
			},
			args: args{
				id: int64(10),
			},
			want: model.Category{
				Template: model.Template{
					ID: uint(10),
				},
				Name: "alat tulis",
			},
		},
		{
			name: "Error-Param",
			mock: func(args args) {
			},
			args: args{
				id: int64(0),
			},
			want:    model.Category{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			got, err := category.DeleteCategory(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
