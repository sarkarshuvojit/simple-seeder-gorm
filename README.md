# simple seeder gorm

Simple framework to structure seeders for your project.

Inspired from https://pkg.go.dev/github.com/sarkarshuvojit/simple-seeder-gorm@v1.0.1/pkg/seeder

## Example Usage

### Declare seeder

```go
// user_seeder.go
var UserSeeder = seeder.Seed{
	Name: "Create users",
	Skip: false,
	Run: func(db *gorm.DB) error {
		for i := 1; i <= 100; i++ {
			user := &models.User{
				Username: fmt.Sprintf("User-%d", i),
				Password: fmt.Sprintf("Pass-%d", i),
			}
			if result := db.Create(user); result.Error != nil {
				return result.Error
			}
		}
		return nil
	},
}
```
### Run it from an entrypoint in your service

```go
// cmd/seed/seed.go
func main() {
	gorm, err := initGorm()
	if err != nil {
		log.Fatal(err)
	}

	if err = seeder.RunSeeder(gorm, []seeder.Seed{UserSeeder}); err != nil {
		slog.Error("Failed to run seeder", "err", err)
	}
}

```
