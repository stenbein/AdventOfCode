package Day15

import scala.io.Source
import scala.util.Using

object Main extends App {

  val filename = "resources/input"
  //val filename = "resources/testInput"

  val parse = raw"(\w+): capacity ([-\d]+), durability ([-\d]+), flavor ([-\d]+), texture ([-\d]+), calories ([-\d]+)".r

  class Ingredient(val name: String, val capacity: Int, val durability: Int, val flavor: Int, val texture: Int, val calories: Int) {
    def rep {println(name, capacity, durability, flavor, texture, calories)}
  }

  class Recipe(val ingredients: Map[Ingredient, Int]) {

    def scoreRecipe(): Int = {

      var capacity: Int = 0
      var durability: Int = 0
      var flavor: Int = 0
      var texture: Int = 0

      for (ingredient <- ingredients.keys) {

        capacity += ingredient.capacity * ingredients(ingredient)
        durability += ingredient.durability * ingredients(ingredient)
        flavor += ingredient.flavor * ingredients(ingredient)
        texture += ingredient.texture * ingredients(ingredient)

      }

      if (capacity < 0 || durability < 0 || flavor < 0 || texture < 0) {return 0}

      capacity * durability * flavor * texture
    }

    def scoreRecipeCals(target: Int): Int = {

      var capacity: Int = 0
      var durability: Int = 0
      var flavor: Int = 0
      var texture: Int = 0
      var calories: Int = 0

      for (ingredient <- ingredients.keys) {

        capacity += ingredient.capacity * ingredients(ingredient)
        durability += ingredient.durability * ingredients(ingredient)
        flavor += ingredient.flavor * ingredients(ingredient)
        texture += ingredient.texture * ingredients(ingredient)
        calories += ingredient.calories * ingredients(ingredient)

      }

      if (capacity < 0 || durability < 0 || flavor < 0 || texture < 0 || calories != target) {return 0}

      capacity * durability * flavor * texture
    }

  }

  val ingredients = Using(Source.fromFile(filename)){
    file => file.getLines().toList.map({
      case parse(name, capacity, durability, flavor, texture, calories)
        => new Ingredient(name, capacity.toInt, durability.toInt, flavor.toInt, texture.toInt, calories.toInt)
    })
  }.get

  def recipe_book(s: List[Ingredient], limit: Int): List[Recipe] = {

    val amounts = s.flatMap(_ => 0 to limit)
      .combinations(s.size)
      .filter(_.sum == 100)
      .flatMap(_.permutations)

    amounts.map(s.zip(_).toMap).map(new Recipe(_)).toList

  }


  def part1() {

    val recipes = recipe_book(ingredients, 100)

    println(recipes.map(_.scoreRecipe()).max)

  }

  def part2() {

    val recipes = recipe_book(ingredients, 100)

    println(recipes.map(_.scoreRecipeCals(500)).max)

  }

  part1()
  part2()

}
