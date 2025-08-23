package repo

// type DataRepo interface {
// 	db.MainRepo[models.Data]
// }

// type datarepo struct {
// 	db.MainRepoImpl[models.Data]
// }

// func NewDataRepo(i *do.Injector) (DataRepo, error) {
// 	return &datarepo{
// 		MainRepoImpl: db.MainRepoImpl[models.Data]{
// 			Db:       do.MustInvoke[*mongo.Client](i),
// 			CollName: "data",
// 		},
// 	}, nil
// }
