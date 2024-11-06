package executors

import (
	"github.com/Jguer/go-alpm/v2"
)

func ReadInstalledPackages(root string, databasePath string) ([]string, []string, error) {
	dbHandle, err := alpm.Initialize(root, databasePath)
	if err != nil {
		return nil, nil, err
	}
	defer dbHandle.Release()

	database, err := dbHandle.LocalDB()
	if err != nil {
		return nil, nil, err
	}

	pkgList := []string{}
	versionList := []string{}

	for _, pkg := range database.PkgCache().Slice() {
		pkgList = append(pkgList, pkg.Name())
		versionList = append(versionList, pkg.Version())
	}

	return pkgList, versionList, nil
}
