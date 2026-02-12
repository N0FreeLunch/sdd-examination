package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"examination/cmd/seeder/internal/seeds"
	"examination/internal/ent"

	entsql "entgo.io/ent/dialect/sql"
	_ "modernc.org/sqlite"
)

func main() {
	// 1. Environment Check
	if os.Getenv("APP_ENV") == "production" {
		log.Fatal("Cannot run seeder in production environment")
	}

	// 2. Parse Flags
	defaultDSN := "file:examination.db?cache=shared&_pragma=foreign_keys(1)"
	seedName := flag.String("seed", "all", "Name of the seeder to run (e.g., exam_preview, all)")
	clean := flag.Bool("clean", false, "Clean existing data before seeding")
	dsnFlag := flag.String("dsn", defaultDSN, "Database source name")
	flag.Parse()

	// Resolve DSN: Flag > Env (DB_PATH) > Default
	dsn := *dsnFlag
	if dsn == defaultDSN {
		if dbPath := os.Getenv("DB_PATH"); dbPath != "" {
			dsn = fmt.Sprintf("file:%s?cache=shared&_pragma=foreign_keys(1)", dbPath)
		}
	}

	// 3. Connect DB
	// Use database/sql directly to support "sqlite" driver (modernc.org/sqlite)
	// while telling Ent to use "sqlite3" dialect.
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	drv := entsql.OpenDB("sqlite3", db)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	ctx := context.Background()

	// 4. Debug: Auto-migrate schema if needed (for local dev convenience)
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// 5. Clean
	if *clean {
		log.Println("Cleaning data...")
		// Simple clean for now: Delete Exams (cascades usually, or manual delete)
		// Since we don't have cascade delete configured in ent explicitly without edges or DB support,
		// we might need to be careful. But sqlite _fk=1 should handle it if schema has ON DELETE CASCADE.
		// If not, we delete manually.
		client.Choice.Delete().Exec(ctx)
		client.ProblemTranslation.Delete().Exec(ctx)
		client.Problem.Delete().Exec(ctx)
		client.Unit.Delete().Exec(ctx)
		client.Section.Delete().Exec(ctx)
		client.Exam.Delete().Exec(ctx)
	}

	// 6. Seed
	log.Printf("Running seeder: %s\n", *seedName)
	switch *seedName {
	case "exam_preview":
		if err := seeds.SeedExamPreview(ctx, client); err != nil {
			log.Fatalf("Failed to seed exam_preview: %v", err)
		}
	case "all":
		// useful for future
		if err := seeds.SeedExamPreview(ctx, client); err != nil {
			log.Fatalf("Failed to seed exam_preview: %v", err)
		}
	default:
		log.Fatalf("Unknown seed name: %s", *seedName)
	}

	log.Println("Seeding completed successfully.")
}
