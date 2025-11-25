CREATE TABLE IF NOT EXISTS movies (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),

    title text NOT NULL,
    description text NOT NULL,
    year integer NOT NULL,
    runtime integer NOT NULL,
    genres text[] NOT NULL,
    languages text[] NOT NULL,

    version integer NOT NULL DEFAULT 1,

    CONSTRAINT movies_year_check 
        CHECK (year BETWEEN 1888 AND date_part('year', now())),

    CONSTRAINT movies_runtime_check 
        CHECK (runtime >= 0),

    CONSTRAINT movies_genres_check 
        CHECK (array_length(genres, 1) BETWEEN 1 AND 5),

    CONSTRAINT movies_languages_check 
        CHECK (array_length(languages, 1) BETWEEN 1 AND 10)
);
