class Config {
  constructor(config) { // HL
    this.str = config.str
    this.i = config.i
  }

  string() {
    return `Config [${this.str}:${this.i}]`
  }
}
// START OMIT
console.log(new Config({ str: 'this is a test' }).string())
console.log(new Config({ str: 'this is a test', i: 8 }).string())
// END OMIT
