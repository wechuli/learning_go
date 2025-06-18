async function doNothing(name: string, year: number): Promise<string> {
  return "nothing";
}
var stuff = doNothing("test", 35);

stuff.then((result) => {
  console.log(result);
});

interface IPerson {
  name: string;
  age: number;
}

class Person {
  constructor(public name: string, public age: number) {
    this.name = name;
    this.age = age;
  }

  async getAge(): Promise<number> {
    return this.age;
  }
}

const cities = ["New York", "Los Angeles", "Chicago", "Houston", "Phoenix"];



const books = {
  title: "The Great Gatsby",
};

const john = new Person("John", 35);

john.getAge().then((age) => {
  console.log(age);
});

// make api call to example.com

fetch("https://jsonplaceholder.typicode.com/posts")
  .then((response) => {
    return response.json();
  })
  .then((data) => {
    console.log(data);
  });
