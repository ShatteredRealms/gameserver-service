package repository_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
)

var _ = Describe("Dimension", func() {
	var repo repository.DimensionRepository
	var dimension *game.Dimension
	BeforeEach(func() {
		repo = repository.NewPgxDimensionRepository(migrater)
		Expect(repo).NotTo(BeNil())
		dimension = &game.Dimension{
			Name:     faker.Username(),
			Location: faker.Username(),
			Version:  faker.Username(),
			Maps:     []*game.Map{},
		}
	})

	Describe("CreateDimension", func() {
		for i := 0; i < 10; i++ {
			It(fmt.Sprintf("should create a dimension with %d maps specified", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				createDimensionsWithMaps(ctx, repo, mapRepo, i)
			})
		}
		When("given invalid input", func() {
			var err error
			var outD *game.Dimension
			It("should error given a nil dimension", func(ctx SpecContext) {
				outD, err = repo.CreateDimension(ctx, nil)
			})
			It("should error given an empty name", func(ctx SpecContext) {
				dimension.Name = ""
				outD, err = repo.CreateDimension(ctx, dimension)
			})
			It("should error given an empty location", func(ctx SpecContext) {
				dimension.Location = ""
				outD, err = repo.CreateDimension(ctx, dimension)
			})
			It("should error given an empty version", func(ctx SpecContext) {
				dimension.Version = ""
				outD, err = repo.CreateDimension(ctx, dimension)
			})
			It("should error given a non-existing map", func(ctx SpecContext) {
				dimension.Maps = []*game.Map{{Id: uuid.New()}}
				outD, err = repo.CreateDimension(ctx, dimension)
			})
			It("should error if the dimension name is taken", func(ctx SpecContext) {
				outD, err = repo.CreateDimension(ctx, dimension)
				Expect(err).NotTo(HaveOccurred())
				Expect(outD).NotTo(BeNil())
				outD, err = repo.CreateDimension(ctx, dimension)

			})
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outD).To(BeNil())
			})
		})
	})

	Describe("DeleteDimension", func() {
		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should delete a dimension with %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, maps := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				outD, err := repo.DeleteDimension(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				expectEquals(outD, dim, true)
				Expect(outD.Maps).To(HaveLen(i))
				for _, outM := range maps {
					Expect(outD.Maps).To(ContainElement(outM))
				}
			})
		}
		It("should error not found given a random id", func(ctx SpecContext) {
			id := uuid.New()
			outD, err := repo.DeleteDimension(ctx, &id)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outD).To(BeNil())
		})
		It("should error given a nil id", func(ctx SpecContext) {
			outD, err := repo.DeleteDimension(ctx, nil)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outD).To(BeNil())
		})
	})

	Describe("GetDeletedDimensions", func() {

		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should return a list of deleted dimensions even if some have %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				deletedD, err := repo.DeleteDimension(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(deletedD).NotTo(BeNil())
				expectEquals(deletedD, dim, true)

				dimensions, err := repo.GetDeletedDimensions(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(dimensions).To(ContainElement(deletedD))
			})
		}
	})

	Describe("GetDimensionById", func() {
		It("should error not found given a random id", func(ctx SpecContext) {
			id := uuid.New()
			outD, err := repo.GetDimensionById(ctx, &id)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outD).To(BeNil())
		})
		It("should error given a nil id", func(ctx SpecContext) {
			outD, err := repo.GetDimensionById(ctx, nil)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outD).To(BeNil())
		})
		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should return a dimension with %d maps even if it was deleted", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, maps := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				deletedD, err := repo.DeleteDimension(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(deletedD).NotTo(BeNil())
				expectEquals(deletedD, dim, true)

				outD, err := repo.GetDimensionById(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				expectEquals(outD, dim, true)
				Expect(outD.Maps).To(HaveLen(i))
				for _, outM := range maps {
					Expect(outD.Maps).To(ContainElement(outM))
				}
			})
		}
		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should return a dimension by id with %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, maps := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				outD, err := repo.GetDimensionById(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				expectEquals(outD, dim, true)
				Expect(outD.Maps).To(HaveLen(i))
				for _, outM := range maps {
					Expect(outD.Maps).To(ContainElement(outM))
				}
			})
		}
	})

	Describe("GetDimensions", func() {
		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should return a list of dimensions even if some have %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				dimensions, err := repo.GetDimensions(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(dimensions).To(ContainElement(dim))
			})
		}

		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should not return deleted dimensions that have %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)
				deletedD, err := repo.DeleteDimension(ctx, &dim.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(deletedD).NotTo(BeNil())

				dimensions, err := repo.GetDimensions(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(dimensions).NotTo(ContainElement(deletedD))
			})
		}
	})

	Describe("UpdateDimension", func() {
		When("given invalid input", func() {
			var err error
			var outD *game.Dimension
			It("should error given a nil dimension", func(ctx SpecContext) {
				outD, err = repo.UpdateDimension(ctx, nil)
			})
			It("should error given an empty name", func(ctx SpecContext) {
				dimension.Name = ""
				outD, err = repo.UpdateDimension(ctx, dimension)
			})
			It("should error given an empty location", func(ctx SpecContext) {
				dimension.Location = ""
				outD, err = repo.UpdateDimension(ctx, dimension)
			})
			It("should error given an empty version", func(ctx SpecContext) {
				dimension.Version = ""
				outD, err = repo.UpdateDimension(ctx, dimension)
			})
			It("should error given a non-existing map", func(ctx SpecContext) {
				dimension.Maps = []*game.Map{{Id: uuid.New()}}
				outD, err = repo.UpdateDimension(ctx, dimension)
			})
			AfterEach(func() {
				Expect(err).To(HaveOccurred())
				Expect(outD).To(BeNil())
			})
		})
		for i := 0; i < 3; i++ {
			It(fmt.Sprintf("should update dimension details with %d maps", i), func(ctx SpecContext) {
				mapRepo := repository.NewPgxMapRepository(migrater)
				Expect(repo).NotTo(BeNil())
				dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)

				deletedAt := time.Now().Add(time.Duration(rand.IntN(1000)) * time.Hour)
				updatedDim := &game.Dimension{
					Id:        dim.Id,
					Name:      faker.Username(),
					Location:  faker.Username(),
					Version:   faker.Username(),
					Maps:      dim.Maps,
					UpdatedAt: time.Now().Add(time.Duration(rand.IntN(1000)) * time.Hour),
					CreatedAt: time.Now().Add(time.Duration(rand.IntN(1000)) * time.Hour),
					DeletedAt: &deletedAt,
				}

				outD, err := repo.UpdateDimension(ctx, updatedDim)
				Expect(err).NotTo(HaveOccurred())
				expectEquals(outD, updatedDim, true)
				Expect(outD.UpdatedAt).NotTo(Equal(updatedDim.UpdatedAt), "should ignore updated at")
				Expect(outD.CreatedAt).NotTo(Equal(updatedDim.CreatedAt), "should ignore created at")
				Expect(outD.DeletedAt).To(BeNil(), "should ignore deleted at")
				Expect(outD.CreatedAt).To(Equal(dim.CreatedAt), "should keep created at")
				Expect(outD.DeletedAt).To(Equal(dim.DeletedAt), "should keep deleted at")
				Expect(outD.UpdatedAt).To(BeTemporally(">", dim.UpdatedAt), "should update updated at")
				Expect(outD.Maps).To(HaveLen(i))
			})
		}
		for i := 0; i < 2; i++ {
			for j := 1; j < 3; j++ {
				It(fmt.Sprintf("should update a dimension going from %d maps to %d", i, i+j), func(ctx SpecContext) {
					mapRepo := repository.NewPgxMapRepository(migrater)
					Expect(repo).NotTo(BeNil())
					dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)
					for k := 0; k < j; k++ {
						m, err := mapRepo.CreateMap(ctx, &game.Map{
							Name:    faker.Username(),
							MapPath: faker.Username(),
						})
						Expect(err).NotTo(HaveOccurred())
						Expect(m).NotTo(BeNil())
						dim.Maps = append(dim.Maps, m)
					}

					outD, err := repo.UpdateDimension(ctx, dim)
					Expect(err).NotTo(HaveOccurred())
					expectEquals(outD, dim, true)
					Expect(outD.Maps).To(HaveLen(i + j))
				})
			}
		}
		for i := 2; i < 4; i++ {
			for j := 1; j < 3; j++ {
				It(fmt.Sprintf("should update a dimension going from %d maps to %d", i, i-j), func(ctx SpecContext) {
					mapRepo := repository.NewPgxMapRepository(migrater)
					Expect(repo).NotTo(BeNil())
					dim, _ := createDimensionsWithMaps(ctx, repo, mapRepo, i)
					GinkgoWriter.Printf("dim maps before: %+v\n", dim.Maps)
					if i-j == 0 {
						dim.Maps = []*game.Map{}
					} else {
						dim.Maps = dim.Maps[:len(dim.Maps)-j]
					}
					GinkgoWriter.Printf("dim maps after: %+v\n", dim.Maps)

					outD, err := repo.UpdateDimension(ctx, dim)
					GinkgoWriter.Printf("dim maps out: %+v\n", outD.Maps)
					Expect(err).NotTo(HaveOccurred())
					expectEquals(outD, dim, true)
					Expect(outD.Maps).To(HaveLen(i - j))
				})
			}
		}
	})
})

func expectEquals(new, original *game.Dimension, checkId bool) {
	Expect(new).NotTo(BeNil(), "new dimension is nil")
	if checkId {
		Expect(new.Id).To(Equal(original.Id), "id does not match")
	}
	Expect(new.Name).To(Equal(original.Name), "name does not match")
	Expect(new.Location).To(Equal(original.Location), "location does not match")
	Expect(new.Version).To(Equal(original.Version), "version does not match")
}

func createDimensionsWithMaps(ctx context.Context, dRepo repository.DimensionRepository, mRepo repository.MapRepository, i int) (*game.Dimension, []*game.Map) {
	dimension := &game.Dimension{
		Name:     faker.Username(),
		Location: faker.Username(),
		Version:  faker.Username(),
	}
	maps := make([]*game.Map, i)
	for j := 0; j < i; j++ {
		m := &game.Map{
			Name:    faker.Username(),
			MapPath: faker.Username(),
		}
		outM, err := mRepo.CreateMap(ctx, m)
		Expect(err).NotTo(HaveOccurred())
		Expect(outM).NotTo(BeNil())
		maps[j] = outM
	}
	dimension.Maps = maps
	d, err := dRepo.CreateDimension(ctx, dimension)
	Expect(err).NotTo(HaveOccurred())
	expectEquals(d, dimension, false)

	Expect(d.Maps).To(HaveLen(i))
	for _, outM := range maps {
		Expect(d.Maps).To(ContainElement(outM))
	}

	return d, maps
}
