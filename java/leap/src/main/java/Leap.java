class Leap {

    boolean isLeapYear(int year) {
        boolean isDivisibleBy4 = year % 4 == 0;
        boolean isDivisibleBy100 = year % 100 == 0;
        boolean isDivisibleBy400 = year % 400 == 0;
        return isDivisibleBy4 && (!isDivisibleBy100 || isDivisibleBy400);
    }

}
