package models

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"meli_test/configs/database"
	"time"
)

// Mutant struct
type Mutant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Dna       string             `json:"dna"`
	IsMutant  bool               `json:"is_mutant"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

// Set Mutant collection and ctx instance
var collection = database.GetCollections("mutants")
var ctx = context.Background()

// IsMutant check if a dna corresponds to the Mutant
func IsMutant(dna []string) bool {
	var matrix Matrix
	matrix.Create(dna)

	var sequence Sequence
	sequence.CheckIsMutant(matrix)
	save(dna, sequence.IsMutant)
	return sequence.IsMutant
}

// FindAll fetch all mutants stored.
func (m Mutant) FindAll() ([]Mutant, error) {
	var mutants []Mutant
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	for cur.Next(ctx) {
		var mutant Mutant
		err = cur.Decode(&mutant)
		mutants = append(mutants, mutant)
	}
	return mutants, err
}

// FindByDna fetch a Mutant by DNA given
func (m Mutant) FindByDna() *Mutant {
	var mutant Mutant
	filter := bson.D{primitive.E{Key: "dna", Value: m.Dna}}

	if err := collection.FindOne(ctx, filter).Decode(&mutant); err != nil {
		return nil
	}
	return &mutant
}

// Create a new Mutant
func (m Mutant) Create() {
	_, err := collection.InsertOne(ctx, &m)
	if err != nil {
		log.Fatalf("Error creating a Mutant %v", err)
	}
}

// save a mutant
func save(dna []string, isMutant bool) {
	dnaString, _ := json.Marshal(dna)
	var mutant = Mutant{
		Dna:       string(dnaString),
		IsMutant:  isMutant,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data := mutant.FindByDna()
	if data != nil {
		return
	}

	mutant.Create()
}
