package migrations

import (
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"gorm.io/gorm"
	"sort"
	"sync"
)

type Migration struct {
	Name          string
	Up            func(db *gorm.DB) error
	Down          func(db *gorm.DB) error
	MigrationType MigrationType
}

type MigrationType uint64

const (
	MigrationDown MigrationType = iota
	MigrationUp
)

var migrationMap = make(map[string]*Migration)
var mu sync.RWMutex

func RegisterMigration(name string, upFunc func(db *gorm.DB) error, downFunc func(db *gorm.DB) error, migrationType MigrationType) {
	mu.Lock()
	defer mu.Unlock()

	migrationMap[name] = &Migration{
		Name:          name,
		Up:            upFunc,
		Down:          downFunc,
		MigrationType: migrationType,
	}
}

func GetMigration(name string) *Migration {
	mu.RLock()
	defer mu.RUnlock()
	return migrationMap[name]
}

func RunMigrations() error {
	mu.RLock()
	defer mu.RUnlock()
	var err error

	var sortedNames []string
	for name := range migrationMap {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	for _, name := range sortedNames {
		migration := GetMigration(name)
		if migration.MigrationType == MigrationUp {
			err = migration.Up(configs.DbClient)
		} else if migration.MigrationType == MigrationDown {
			err = migration.Down(configs.DbClient)
		}

		if err != nil {
			return err
		}
		fmt.Printf("Migration %s completed successfully\n", migration.Name)
	}

	fmt.Println("All migrations completed successfully")
	return nil
}

func RollbackMigrations() error {
    mu.RLock()
    defer mu.RUnlock()
    var err error

    var sortedNames []string
    for name := range migrationMap {
        sortedNames = append(sortedNames, name)
    }
    sort.Sort(sort.Reverse(sort.StringSlice(sortedNames)))

    for _, name := range sortedNames {
        migration := GetMigration(name)
        if migration.MigrationType == MigrationDown {
            err = migration.Up(configs.DbClient)
        } else if migration.MigrationType == MigrationUp {
            err = migration.Down(configs.DbClient)
        }

        if err != nil {
            return err
        }
        fmt.Printf("Rollback of Migration %s completed successfully\n", migration.Name)
    }

    fmt.Println("All migrations rolled back successfully")
    return nil
}