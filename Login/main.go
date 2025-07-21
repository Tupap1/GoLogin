package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2" 
	"github.com/jackc/pgx/v5"     
	"github.com/joho/godotenv"    