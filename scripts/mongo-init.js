db.createUser({
    user: "admin",
    pwd: "password",  // Replace with your desired password
    roles: [{ role: "root", db: "admin" }]
});

db.createUser({
    user: "dev",
    pwd: "passdev",
    roles: ["root"]
});

// Create DB and add data
db = db.getSiblingDB('coffee-chooser');

db.coffee_types.insertMany([
    {
        "type": "Arara",
        "sweetness": "Very Sweet",
        "strength": "Moderately Strong",
        "flavorNotes": "Fruity, Floral",
        "body": "Full",
        "description": "Arara is a Brazilian coffee variety known for its vibrant fruity and floral profile, with a balanced sweetness and a full-bodied taste."
    },
    {
        "type": "Geisha",
        "sweetness": "Moderately Sweet",
        "strength": "Mild",
        "flavorNotes": "Citrus, Jasmine",
        "body": "Light",
        "description": "Geisha is a renowned coffee variety praised for its delicate sweetness, mild strength, and light body, with complex flavor notes of citrus and jasmine."
    },
    {
        "type": "Mundo Novo",
        "sweetness": "Sweet",
        "strength": "Strong",
        "flavorNotes": "Chocolate, Nutty",
        "body": "Full",
        "description": "Mundo Novo is a classic Brazilian coffee variety with a strong profile, full body, and sweet notes of chocolate and nuts."
    },
    {
        "type": "Catuai",
        "sweetness": "Moderately Sweet",
        "strength": "Moderate",
        "flavorNotes": "Honey, Almond",
        "body": "Medium",
        "description": "Catuai offers a well-balanced flavor with moderate sweetness and strength, complemented by honey and almond notes."
    },
    {
        "type": "Bourbon",
        "sweetness": "Sweet",
        "strength": "Moderately Strong",
        "flavorNotes": "Caramel, Berry",
        "body": "Full",
        "description": "Bourbon coffee is celebrated for its full body, moderately strong profile, and sweet caramel and berry notes."
    },
    {
        "type": "Icatu",
        "sweetness": "Very Sweet",
        "strength": "Strong",
        "flavorNotes": "Spice, Chocolate",
        "body": "Full",
        "description": "Icatu is known for its very sweet flavor profile, strong strength, and rich notes of spice and chocolate."
    },
    {
        "type": "Typica",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Sweet, Clean",
        "body": "Medium",
        "description": "Typica is one of the oldest Arabica varieties, known for its clean, sweet taste and balanced acidity, making it a classic choice."
    },
    {
        "type": "SL28",
        "sweetness": "Moderately Sweet",
        "strength": "Strong",
        "flavorNotes": "Fruity, Bright Acidity",
        "body": "Medium",
        "description": "SL28, originating from Kenya, is celebrated for its strong fruit notes and bright acidity, offering a vibrant coffee experience."
    },
    {
        "type": "SL34",
        "sweetness": "Moderately Sweet",
        "strength": "Strong",
        "flavorNotes": "Rich, Fruity",
        "body": "Full",
        "description": "SL34 is similar to SL28 but with a richer body and enhanced disease resistance, delivering a robust and fruity cup."
    },
    {
        "type": "Pacamara",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Fruity, Complex",
        "body": "Full",
        "description": "Pacamara, a hybrid of Pacas and Maragogipe, is known for its large beans and complex flavors, often fruity and sweet."
    },
    {
        "type": "Maragogipe",
        "sweetness": "Sweet",
        "strength": "Mild",
        "flavorNotes": "Mild, Smooth",
        "body": "Light",
        "description": "Also known as 'Elephant Bean' due to its large size, Maragogipe has a mild flavor with a smooth body, perfect for a gentle cup."
    },
    {
        "type": "Pink Bourbon",
        "sweetness": "Very Sweet",
        "strength": "Mild",
        "flavorNotes": "Floral, Fruity",
        "body": "Light",
        "description": "A rare variety from Colombia, Pink Bourbon is known for its floral aroma and fruity, complex flavors with a light body."
    },
    {
        "type": "Java",
        "sweetness": "Moderately Sweet",
        "strength": "Strong",
        "flavorNotes": "Earthy, Spicy",
        "body": "Full",
        "description": "Originally from Indonesia, Java coffee is famous for its earthy, spicy flavors, offering a unique and full-bodied experience."
    },
    {
        "type": "Mokka",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Chocolatey",
        "body": "Medium",
        "description": "Mokka is a small bean with a distinct chocolatey flavor, often used in high-end blends for its rich and smooth taste."
    },
    {
        "type": "Yellow Bourbon",
        "sweetness": "Very Sweet",
        "strength": "Moderately Strong",
        "flavorNotes": "Fruity, Sweet",
        "body": "Full",
        "description": "A mutation of Bourbon with yellow cherries, Yellow Bourbon is known for its very sweet profile and balanced acidity."
    },
    {
        "type": "Caturra",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Bright, Fruity",
        "body": "Medium",
        "description": "Caturra is a natural mutation of Bourbon, popular in Latin America for its bright acidity and fruity flavor profile."
    },
    {
        "type": "Obata",
        "sweetness": "Moderately Sweet",
        "strength": "Moderate",
        "flavorNotes": "Smooth, Sweet",
        "body": "Medium",
        "description": "Obata is a hybrid developed in Brazil, resistant to disease with a smooth, sweet flavor profile, perfect for a balanced cup."
    },
    {
        "type": "Pacas",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Fruity, Mild",
        "body": "Medium",
        "description": "Pacas, a mutation of Bourbon, is known for its balanced cup with a hint of fruitiness, offering a mild and sweet experience."
    },
    {
        "type": "Laurina",
        "sweetness": "Sweet",
        "strength": "Mild",
        "flavorNotes": "Mild, Sweet",
        "body": "Light",
        "description": "Also known as 'Bourbon Pointu,' Laurina is low in caffeine and has a mild, sweet flavor, making it a unique and gentle cup."
    },
    {
        "type": "Sarchimor",
        "sweetness": "Moderately Sweet",
        "strength": "Moderate",
        "flavorNotes": "Chocolatey, Smooth",
        "body": "Medium",
        "description": "Sarchimor is a hybrid known for its disease resistance and balanced flavor profile, often with a hint of chocolate."
    },
    {
        "type": "Villalobos",
        "sweetness": "Sweet",
        "strength": "Moderate",
        "flavorNotes": "Fruity, Sweet",
        "body": "Medium",
        "description": "Villalobos is a lesser-known Costa Rican variety, offering a sweet and fruity flavor profile with a balanced body."
    }
]);

db.producers.insertMany([
    {
        "name": "Fazenda Santa Inês",
        "location": "Minas Gerais, Brazil",
        "producedBeans": ["Bourbon", "Yellow Bourbon", "Catuai"],
        "description": "Located in the Mogiana region, Fazenda Santa Inês is known for producing high-quality Bourbon and Catuai varieties."
    },
    {
        "name": "Fazenda Daterra",
        "location": "Cerrado Mineiro, Brazil",
        "producedBeans": ["Bourbon", "Icatu", "Arara", "Yellow Bourbon"],
        "description": "Daterra is a leading farm in Brazil, known for innovation and sustainability, producing a wide range of beans."
    },
    {
        "name": "Fazenda Samambaia",
        "location": "Sul de Minas, Brazil",
        "producedBeans": ["Bourbon", "Yellow Bourbon", "Mundo Novo"],
        "description": "Fazenda Samambaia is recognized for its innovative practices and high-quality coffee production."
    },
    {
        "name": "Fazenda Ambiental Fortaleza",
        "location": "Mococa, São Paulo, Brazil",
        "producedBeans": ["Bourbon", "Catuai", "Icatu"],
        "description": "A pioneer in sustainable farming, Fazenda Ambiental Fortaleza is known for its premium Bourbon and Catuai beans."
    },
    {
        "name": "Fazenda Esperança",
        "location": "Carmo de Minas, Brazil",
        "producedBeans": ["Yellow Bourbon", "Catuai"],
        "description": "Located in Carmo de Minas, Fazenda Esperança produces exceptional Yellow Bourbon and Catuai varieties."
    },
    {
        "name": "Fazenda Rio Verde",
        "location": "Sul de Minas, Brazil",
        "producedBeans": ["Mundo Novo", "Bourbon"],
        "description": "One of the oldest coffee farms in Brazil, Fazenda Rio Verde is known for its rich Mundo Novo and Bourbon beans."
    },
    {
        "name": "Fazenda São Francisco",
        "location": "Cerrado Mineiro, Brazil",
        "producedBeans": ["Icatu", "Catuai"],
        "description": "Fazenda São Francisco focuses on producing robust Icatu and Catuai beans, with an emphasis on quality."
    },
    {
        "name": "Fazenda Fortaleza",
        "location": "Mococa, São Paulo, Brazil",
        "producedBeans": ["Bourbon", "Mundo Novo"],
        "description": "Fazenda Fortaleza is committed to producing high-quality coffee, particularly Bourbon and Mundo Novo varieties."
    }
]);

db.coffee_brands.insertMany([
    {
        "name": "Orfeu Cafés Especiais",
        "description": "A leading Brazilian specialty coffee brand committed to quality and sustainability, sourcing beans from their own farms in Sul de Minas and Mogiana."
    },
    {
        "name": "Coffee++",
        "description": "A modern specialty coffee brand from Brazil, Coffee++ focuses on delivering high-quality coffee with an emphasis on traceability and direct trade."
    },
    {
        "name": "Santa Monica",
        "description": "Café Santa Monica is one of Brazil's oldest specialty coffee brands, known for its high-quality beans and traditional processing methods."
    },
    {
        "name": "Um Coffee Co",
        "description": "Um Coffee Co is a specialty coffee roaster based in São Paulo, dedicated to sourcing and roasting the finest Brazilian coffees."
    },
    {
        "name": "Café 3 Corações",
        "description": "One of Brazil’s largest coffee brands, offering a variety of blends made from beans sourced across Minas Gerais and São Paulo."
    },
    {
        "name": "Café Santo Grão",
        "description": "A premium coffee brand working closely with farms in Mogiana and Sul de Minas, known for their single-origin coffees."
    },
    {
        "name": "Café Suplicy",
        "description": "A specialty coffee brand based in São Paulo, focused on direct trade with top farms and offering premium coffee options."
    },
    {
        "name": "Café Octavio",
        "description": "A family-owned brand focusing on specialty coffee, producing high-quality coffees from the Alta Mogiana region."
    }
]);

// Indexes for fast search
db.producers.createIndex({name: 1});
db.producers.createIndex({location: 1});
db.producers.createIndex({producedBeans: 1});

db.coffee_types.createIndex({type: 1});
db.coffee_types.createIndex({sweetness: 1, strength: 1, flavorNotes: 1, body: 1});

db.coffee_brands.createIndex({name: 1});
db.coffee_brands.createIndex({description: "text"});
