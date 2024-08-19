// Switch to the coffee-chooser database to add data
db = db.getSiblingDB('coffee-chooser');

// Insert coffee types
db.coffee_types.insertMany([
    {
        "variety": "arara",
        "sweetness": "very_sweet",
        "strength": "moderately_strong",
        "flavor_notes": "fruity, floral",
        "body": "full",
        "description": "arara is a brazilian coffee variety known for its vibrant fruity and floral profile, with a balanced sweetness and a full-bodied taste.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "geisha",
        "sweetness": "moderately_sweet",
        "strength": "mild",
        "flavor_notes": "citrus, jasmine",
        "body": "light",
        "description": "geisha is a renowned coffee variety praised for its delicate sweetness, mild strength, and light body, with complex flavor notes of citrus and jasmine.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "mundo_novo",
        "sweetness": "sweet",
        "strength": "strong",
        "flavor_notes": "chocolate, nutty",
        "body": "full",
        "description": "mundo novo is a classic brazilian coffee variety with a strong profile, full body, and sweet notes of chocolate and nuts.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "catuai",
        "sweetness": "moderately_sweet",
        "strength": "moderate",
        "flavor_notes": "honey, almond",
        "body": "medium",
        "description": "catuai offers a well-balanced flavor with moderate sweetness and strength, complemented by honey and almond notes.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "bourbon",
        "sweetness": "sweet",
        "strength": "moderately_strong",
        "flavor_notes": "caramel, berry",
        "body": "full",
        "description": "bourbon coffee is celebrated for its full body, moderately strong profile, and sweet caramel and berry notes.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "icatu",
        "sweetness": "very_sweet",
        "strength": "strong",
        "flavor_notes": "spice, chocolate",
        "body": "full",
        "description": "icatu is known for its very sweet flavor profile, strong strength, and rich notes of spice and chocolate.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "typica",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "sweet, clean",
        "body": "medium",
        "description": "typica is one of the oldest arabica varieties, known for its clean, sweet taste and balanced acidity, making it a classic choice.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "sl28",
        "sweetness": "moderately_sweet",
        "strength": "strong",
        "flavor_notes": "fruity, bright_acidity",
        "body": "medium",
        "description": "sl28, originating from kenya, is celebrated for its strong fruit notes and bright acidity, offering a vibrant coffee experience.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "sl34",
        "sweetness": "moderately_sweet",
        "strength": "strong",
        "flavor_notes": "rich, fruity",
        "body": "full",
        "description": "sl34 is similar to sl28 but with a richer body and enhanced disease resistance, delivering a robust and fruity cup.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "pacamara",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "fruity, complex",
        "body": "full",
        "description": "pacamara, a hybrid of pacas and maragogipe, is known for its large beans and complex flavors, often fruity and sweet.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "maragogipe",
        "sweetness": "sweet",
        "strength": "mild",
        "flavor_notes": "mild, smooth",
        "body": "light",
        "description": "also known as 'elephant bean' due to its large size, maragogipe has a mild flavor with a smooth body, perfect for a gentle cup.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "pink_bourbon",
        "sweetness": "very_sweet",
        "strength": "mild",
        "flavor_notes": "floral, fruity",
        "body": "light",
        "description": "a rare variety from colombia, pink bourbon is known for its floral aroma and fruity, complex flavors with a light body.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "java",
        "sweetness": "moderately_sweet",
        "strength": "strong",
        "flavor_notes": "earthy, spicy",
        "body": "full",
        "description": "originally from indonesia, java coffee is famous for its earthy, spicy flavors, offering a unique and full-bodied experience.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "mokka",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "chocolatey",
        "body": "medium",
        "description": "mokka is a small bean with a distinct chocolatey flavor, often used in high-end blends for its rich and smooth taste.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "yellow_bourbon",
        "sweetness": "very_sweet",
        "strength": "moderately_strong",
        "flavor_notes": "fruity, sweet",
        "body": "full",
        "description": "a mutation of bourbon with yellow cherries, yellow bourbon is known for its very sweet profile and balanced acidity.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "caturra",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "bright, fruity",
        "body": "medium",
        "description": "caturra is a natural mutation of bourbon, popular in latin america for its bright acidity and fruity flavor profile.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "obata",
        "sweetness": "moderately_sweet",
        "strength": "moderate",
        "flavor_notes": "smooth, sweet",
        "body": "medium",
        "description": "obata is a hybrid developed in brazil, resistant to disease with a smooth, sweet flavor profile, perfect for a balanced cup.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "pacas",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "fruity, mild",
        "body": "medium",
        "description": "pacas, a mutation of bourbon, is known for its balanced cup with a hint of fruitiness, offering a mild and sweet experience.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "laurina",
        "sweetness": "sweet",
        "strength": "mild",
        "flavor_notes": "mild, sweet",
        "body": "light",
        "description": "also known as 'bourbon pointu,' laurina is low in caffeine and has a mild, sweet flavor, making it a unique and gentle cup.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "sarchimor",
        "sweetness": "moderately_sweet",
        "strength": "moderate",
        "flavor_notes": "chocolatey, smooth",
        "body": "medium",
        "description": "sarchimor is a hybrid known for its disease resistance and balanced flavor profile, often with a hint of chocolate.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "villalobos",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "fruity, sweet",
        "body": "medium",
        "description": "villalobos is a lesser-known costa rican variety, offering a sweet and fruity flavor profile with a balanced body.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "acaia",
        "sweetness": "moderately_sweet",
        "strength": "moderate",
        "flavor_notes": "nutty, sweet",
        "body": "medium",
        "description": "acaia is known for its large beans, mild acidity, and a sweet, nutty flavor profile, often used in blends or as a single-origin coffee.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "tupi",
        "sweetness": "sweet",
        "strength": "strong",
        "flavor_notes": "chocolate, sweet",
        "body": "full",
        "description": "tupi is a hybrid variety known for its resistance to rust and its sweet, chocolatey flavor profile.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "topazio",
        "sweetness": "moderately_sweet",
        "strength": "moderate",
        "flavor_notes": "fruity, floral",
        "body": "medium",
        "description": "topazio is known for its balanced profile with moderate acidity and body, featuring fruity and floral notes.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "sabiá",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "sweet, mild",
        "body": "medium",
        "description": "sabiá is a newer variety recognized for its resistance to rust and nematodes, with a balanced flavor profile.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "rubí",
        "sweetness": "sweet",
        "strength": "strong",
        "flavor_notes": "chocolate, nutty",
        "body": "full",
        "description": "rubí is known for its strong body and lower acidity, often exhibiting chocolate and nutty flavors.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "paraiso",
        "sweetness": "very_sweet",
        "strength": "moderate",
        "flavor_notes": "floral, fruity",
        "body": "full",
        "description": "paraiso is appreciated for its intense sweetness, floral aroma, and complex fruity flavors, making it popular among specialty coffee growers.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "oeiras",
        "sweetness": "sweet",
        "strength": "moderate",
        "flavor_notes": "chocolate, nutty",
        "body": "medium",
        "description": "oeiras is known for its smooth body and balanced acidity, with notes of chocolate and nuts, making it a popular variety in brazil.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "variety": "bourbon amarelo",
        "sweetness": "sweet",
        "strength": "moderately_strong",
        "flavor_notes": "citrus, fruity",
        "body": "full",
        "description": "bourbon amarelo is a mutation of the bourbon variety, known for its bright acidity, sweetness, and complex fruit flavors.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    }
]);

// Insert coffee producers
db.producers.insertMany([
    {
        "name": "fazenda santa inês",
        "location": "minas gerais, brazil",
        "producedBeans": ["bourbon", "yellow bourbon", "catuai"],
        "description": "located in the mogiana region, fazenda santa inês is known for producing high-quality bourbon and catuai varieties.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda daterra",
        "location": "cerrado mineiro, brazil",
        "producedBeans": ["bourbon", "icatu", "arara", "yellow bourbon"],
        "description": "daterra is a leading farm in brazil, known for innovation and sustainability, producing a wide range of beans.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda samambaia",
        "location": "sul de minas, brazil",
        "producedBeans": ["bourbon", "yellow bourbon", "mundo novo"],
        "description": "fazenda samambaia is recognized for its innovative practices and high-quality coffee production.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda ambiental fortaleza",
        "location": "mococa, são paulo, brazil",
        "producedBeans": ["bourbon", "catuai", "icatu"],
        "description": "a pioneer in sustainable farming, fazenda ambiental fortaleza is known for its premium bourbon and catuai beans.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda esperança",
        "location": "carmo de minas, brazil",
        "producedBeans": ["yellow bourbon", "catuai"],
        "description": "located in carmo de minas, fazenda esperança produces exceptional yellow bourbon and catuai varieties.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda rio verde",
        "location": "sul de minas, brazil",
        "producedBeans": ["mundo novo", "bourbon"],
        "description": "one of the oldest coffee farms in brazil, fazenda rio verde is known for its rich mundo novo and bourbon beans.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda são francisco",
        "location": "cerrado mineiro, brazil",
        "producedBeans": ["icatu", "catuai"],
        "description": "fazenda são francisco focuses on producing robust icatu and catuai beans, with an emphasis on quality.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda fortaleza",
        "location": "mococa, são paulo, brazil",
        "producedBeans": ["bourbon", "mundo novo"],
        "description": "fazenda fortaleza is committed to producing high-quality coffee, particularly bourbon and mundo novo varieties.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda boa vista",
        "location": "espirito santo, brazil",
        "producedBeans": ["acaia", "catuai", "bourbon amarelo"],
        "description": "fazenda boa vista is known for producing high-quality acaia and bourbon amarelo varieties in the mountains of espirito santo.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda das mangabeiras",
        "location": "bahia, brazil",
        "producedBeans": ["tupi", "topazio", "sarchimor"],
        "description": "located in bahia, fazenda das mangabeiras is a leader in sustainable coffee production, specializing in tupi, topazio, and sarchimor varieties.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda santa maria",
        "location": "sul de minas, brazil",
        "producedBeans": ["rubí", "paraiso", "oeiras"],
        "description": "fazenda santa maria, situated in sul de minas, produces some of the finest rubí, paraiso, and oeiras coffees.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "fazenda das palmeiras",
        "location": "são paulo, brazil",
        "producedBeans": ["sabiá", "caturra", "mundo novo"],
        "description": "fazenda das palmeiras in são paulo focuses on growing high-quality sabiá and caturra, along with the traditional mundo novo.",
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    }
]);


// Insert coffee brands
db.coffee_brands.insertMany([
    {
        "name": "orfeu cafés especiais",
        "description": "a leading brazilian specialty coffee brand committed to quality and sustainability, sourcing beans from their own farms in sul de minas and mogiana.",
        "sold_varieties": ["yellow_bourbon", "catuai", "icatu"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "coffee++",
        "description": "a modern specialty coffee brand from brazil, coffee++ focuses on delivering high-quality coffee with an emphasis on traceability and direct trade.",
        "sold_varieties": ["arara", "bourbon", "catuai"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "santa monica",
        "description": "café santa monica is one of brazil's oldest specialty coffee brands, known for its high-quality beans and traditional processing methods.",
        "sold_varieties": ["mundo_novo", "yellow_bourbon", "typica"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "um coffee co",
        "description": "um coffee co is a specialty coffee roaster based in são paulo, dedicated to sourcing and roasting the finest brazilian coffees.",
        "sold_varieties": ["geisha", "icatu", "mokka"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "café 3 corações",
        "description": "one of brazil’s largest coffee brands, offering a variety of blends made from beans sourced across minas gerais and são paulo.",
        "sold_varieties": ["bourbon", "catuai", "mundo_novo"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "café santo grão",
        "description": "a premium coffee brand working closely with farms in mogiana and sul de minas, known for their single-origin coffees.",
        "sold_varieties": ["pacamara", "laurina", "java"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "café suplicy",
        "description": "a specialty coffee brand based in são paulo, focused on direct trade with top farms and offering premium coffee options.",
        "sold_varieties": ["yellow_bourbon", "sl28", "sl34"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "café octavio",
        "description": "a family-owned brand focusing on specialty coffee, producing high-quality coffees from the alta mogiana region.",
        "sold_varieties": ["obata", "pacas", "sarchimor"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "lucca cafés especiais",
        "description": "lucca cafés especiais is known for its wide range of specialty coffee varieties and a strong focus on quality and education within the brazilian coffee scene.",
        "sold_varieties": ["pink_bourbon", "maragogipe", "villalobos"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "café do sítio",
        "description": "café do sítio sources from small farms across brazil, offering a diverse range of specialty coffees known for their rich and complex flavors.",
        "sold_varieties": ["pacamara", "geisha", "mundo_novo"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "coffeelab",
        "description": "coffeelab, led by isabela raposeiras, is one of brazil's most innovative specialty coffee brands, offering carefully curated and roasted coffees with a strong emphasis on quality and sustainability.",
        "sold_varieties": ["mokka", "pink_bourbon", "laurina"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "moka clube",
        "description": "moka clube is a subscription-based specialty coffee brand that offers unique and limited-edition coffees sourced from small farms across brazil.",
        "sold_varieties": ["obata", "sarchimor", "bourbon"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "torra clara",
        "description": "torra clara is known for its commitment to light roasting, allowing the natural flavors of high-quality beans to shine through, sourced from select farms in brazil.",
        "sold_varieties": ["sl28", "sl34", "typica"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    },
    {
        "name": "academia do café",
        "description": "based in belo horizonte, academia do café is a leading brand in the specialty coffee movement in brazil, offering expertly roasted beans from farms they work closely with.",
        "sold_varieties": ["pacamara", "mundo_novo", "catuai"],
        "created_at": new Date().toISOString(),
        "updated_at": new Date().toISOString()
    }
]);

// Create indexes for faster search and retrieval
db.producers.createIndex({name: 1});
db.producers.createIndex({location: 1});
db.producers.createIndex({producedBeans: 1});
db.producers.createIndex({created_at: 1, updated_at: 1});

db.coffee_types.createIndex({type: 1});
db.coffee_types.createIndex({sweetness: 1, strength: 1, flavor_notes: 1, body: 1});
db.coffee_types.createIndex({created_at: 1, updated_at: 1});

db.coffee_brands.createIndex({name: 1});
db.coffee_brands.createIndex({description: "text"});
db.coffee_brands.createIndex({created_at: 1, updated_at: 1});
