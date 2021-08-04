public class Twofer {
    public String twofer(String name) {
        String msg;
        if (name == null) {
            msg = "One for you, one for me.";
        } else {
            msg = "One for " + name + ", one for me.";
        }
        return msg;
    }
}
