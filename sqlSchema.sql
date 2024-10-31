DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'square_size') THEN
        CREATE TYPE square_size AS ENUM ('Tiny', 'Small', 'Medium', 'Large', 'Huge', 'Gargantuan');
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
        CREATE TYPE mastery_level AS ENUM ('None', 'Trained', 'Expert', 'Master', 'Legend');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ability') THEN
CREATE TYPE ability AS ENUM ('Strength', 'Dexterity', 'Constitution', 'Intelligence', 'Wisdom', 'Charisma');
END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'rarity') THEN
CREATE TYPE rarity AS ENUM ('Common', 'Uncommon', 'Rare', 'Legendary', 'Mythic');
END IF;
END $$;