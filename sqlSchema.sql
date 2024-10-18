DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'square_size') THEN
CREATE TYPE square_size AS ENUM ('Tiny', 'Small', 'Medium', 'Large', 'Huge', 'Gargantuan');
END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'hit_point') THEN
CREATE TYPE hit_point AS ENUM ('Six', 'Eight', 'Ten', 'Twelve');
END IF;
END $$;
