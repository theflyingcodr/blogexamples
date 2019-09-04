package main

import "context"

type Airport struct{
	Name string
	City string
	Country string
	Runways  []Runway
}

type Runway struct{
	Bearing string
	Length int
}

type AirportService interface {
	Create(ctx context.Context, a Airport) (*Airport, error)
}

