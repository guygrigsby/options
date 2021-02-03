class Overload {
  String str;
  int i;
  Overload() { // HL
    this("default", 9);
  }

  Overload(String str) { // HL
    this(str, 9);
  }

  Overload(String str, int i) { // HL
    this.str = str;
    this.i =  i;
  }

  public String toString() {
    return "Overload [" + this.str + ":" + this.i+"]";
  }
  public static void main(String[] args) {
    int i = 4;
    String str = "this is a test";

    Overload c = new Overload();
    System.out.println(c);

    c = new Overload(str);
    System.out.println(c);

    c = new Overload(str, i);
    System.out.println(c);
  }

}
