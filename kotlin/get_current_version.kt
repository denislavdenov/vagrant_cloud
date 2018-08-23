fun main(args: Array<String>) {


print("The box of which user do you want?: ")
val var1 = readLine()!!
print("The url of which box do you want?: ")
val var2 = readLine()!!

println("You want the box of $var1/$var2")

val url = "https://app.vagrantup.com/api/v1/$var1/$var2"

println("Box can be found on $url")

}
