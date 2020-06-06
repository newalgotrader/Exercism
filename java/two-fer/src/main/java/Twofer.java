import java.util.Objects;

public class Twofer {
    public String twofer(String name) {
      if (Objects.isNull(name)) {
        name = "you";
      }

      return "One for " + name + ", one for me.";
    }
}
