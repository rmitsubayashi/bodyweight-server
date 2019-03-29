# API Docs
Responses follow the [JSON API specifications](https://jsonapi.org/)

## Exercise
**id** int

**title** string

**level** int

**image_url** string

**description** string

**[measurement_type](#measurement-type)** string

**category_id** int

**target_sets** Array(Set)

**quantity** int //-1 = infinite

### GET /users/exercises/category_id=??
returns Array(Exercise)

## ExerciseProduct
**title** string

**exercises** Array(Exercise)

**price** int

**sold** boolean

### GET /shop/exercises
returns Array(ExerciseProduct)

### POST /shop/exercises
returns nothing

## User
**id** int

**firebase_uid** string

**points** int

**cat1_level** int

...

**cat6_level** int

### GET /users
returns User

### POST /users

## Log
**id** int

**category_id** int

**date** string //YYYY-MM-DD

**sets** Array([Set](#set)) 

### GET /users/logs
returns Array(Log)

### GET /users/logs/log_id
returns Log

### POST /users/logs
returns Feedback

## Feedback
**comment** string

**comment_highlight_spans** Array([int, int]) //inclusive-exclusive

**previous_points** int

**after_points** int

**level_up** bool

**unlocked_exercises** Array(UnlockedExercise)

**dropped_exercises** Array(Exercise)

## UnlockedExercise
**exercise** Exercise

**other_exercises** Array(Exercise)

## Set

**exercise** Exercise

**set_number** int

**value** int

## Measurement Type
REP, SECONDS
