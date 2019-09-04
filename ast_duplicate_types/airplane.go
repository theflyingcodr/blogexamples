package main

import "context"

type Airplane struct{
	Manufacturer string
	Model string
	Airline string
}

type AirplanceService interface{
	Airplane(ctx context.Context, model string) (*Airplane, error)
}

