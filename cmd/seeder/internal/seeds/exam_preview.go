package seeds

import (
	"context"
	"examination/internal/ent"
	"fmt"
	"time"
)

func SeedExamPreview(ctx context.Context, client *ent.Client) error {
	// 1. Create Exam
	e, err := client.Exam.Create().
		SetTitle("Distributed Systems 101").
		SetDescription("Preview exam for data verification").
		SetTimeLimit(60).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating exam: %w", err)
	}

	// 2. Create Section
	s, err := client.Section.Create().
		SetTitle("Data Consistency").
		SetSeq(1).
		SetExam(e).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating section: %w", err)
	}

	// Helper to create full question stack
	createQuestion := func(seq int, title, content, explanation string, choices []string, correctIdx int, codeBlock string) error {
		// Unit
		u, err := client.Unit.Create().
			SetTitle(title).
			SetSeq(seq).
			SetExam(e).
			SetSection(s).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating unit %d: %w", seq, err)
		}

		// Problem
		p, err := client.Problem.Create().
			SetUnit(u).
			SetType("SOURCE").
			SetCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating problem %d: %w", seq, err)
		}

		// Translation
		finalContent := content
		if codeBlock != "" {
			finalContent += "\n\n" + codeBlock
		}

		pt, err := client.ProblemTranslation.Create().
			SetProblem(p).
			SetLocale("en").
			SetTitle(title).
			SetContent(finalContent).
			SetExplanation(explanation).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating translation %d: %w", seq, err)
		}

		// Choices
		builders := make([]*ent.ChoiceCreate, len(choices))
		for i, c := range choices {
			builders[i] = client.Choice.Create().
				SetProblemTranslation(pt).
				SetContent(c).
				SetSeq(i + 1).
				SetIsCorrect(i == correctIdx)
		}
		if _, err := client.Choice.CreateBulk(builders...).Save(ctx); err != nil {
			return fmt.Errorf("failed creating choices for unit %d: %w", seq, err)
		}

		return nil
	}

	// Unit 1: Simple
	if err := createQuestion(1,
		"CAP Theorem",
		"According to the CAP theorem, which two properties can be satisfied simultaneously in the presence of validation?",
		"In a distributed system with partitions (P), you must choose between Consistency (C) and Availability (A).",
		[]string{"Consistency & Availability", "Availability & Partition Tolerance", "Consistency & Partition Tolerance", "None of the above"},
		2, // CP is common choice for strong consistency systems
		"",
	); err != nil {
		return err
	}

	// Unit 2: Long Text
	longText := `Eventual consistency is a consistency model used in distributed computing to achieve high availability that informally guarantees that, if no new updates are made to a given data item, eventually all accesses to that item will return the last updated value.

Which of the following scenarios BEST describes a system prioritizing Eventual Consistency?`
	if err := createQuestion(2,
		"Eventual Consistency",
		longText,
		"DNS is a classic example where changes take time to propagate, but eventually all resolvers see the new record.",
		[]string{
			"A banking ledger transaction system.",
			"A real-time stock trading engine.",
			"A DNS change propagation system.",
			"A pacemaker control system.",
		},
		2,
		"",
	); err != nil {
		return err
	}

	// Unit 3: Code
	codeBlock := "```go\nfunc main() {\n\tch := make(chan int, 1)\n\tch <- 1\n\tfmt.Println(<-ch)\n}\n```"
	if err := createQuestion(3,
		"Go Channels",
		"What is the output of the following Go code?",
		"The channel is buffered with size 1. The send does not block because there is space. The receive reads the value.",
		[]string{"1", "Deadlock", "Runtime Panic", "Compilation Error"},
		0,
		codeBlock,
	); err != nil {
		return err
	}

	return nil
}
