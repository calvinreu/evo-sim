# Protocol to save

## Main JSON

the main json file contains the relative file paths to evosim of the other required config files it should look like this:

    {
        "configFile" : "path to main config file",
        "mapFile" : "path to map file",
        "creatures" : ["creature file CreatureNr", ...]
    }

## Map file

the map file should contain all the atributes of the current map it should look like

    {
        "Fields" : [
            [ "Field" : {
                "Food" : some value,
                "WaterPercentage: 0 <= value <= 100,
                "Creatures" : [CreatureNr, ...]
            } ],
             ...]
    }

## Creature file

    {
        "Brain" : "path to neural net file",
        "Food" : foodValue,
        "Water" : water value of the creature,
        "Speed" : speed at WaterPercentage 0,
        "WaterSlowdown" : value of slowdown
    }
