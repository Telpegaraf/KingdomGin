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

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tradition') THEN
        CREATE TYPE tradition AS ENUM ('None', 'Arcane', 'Divine', 'Occult', 'Primal');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'school') THEN
        CREATE TYPE school AS ENUM
            ('Abjuration', 'Conjuration', 'Divination', 'Enchantment',
            'Evocation', 'Illusion', 'Necromancy', 'Transmutation');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mastery_level') THEN
        CREATE TYPE mastery_level AS ENUM ('None', 'Train', 'Expert', 'Master', 'Legend');
    END IF;
END $$;