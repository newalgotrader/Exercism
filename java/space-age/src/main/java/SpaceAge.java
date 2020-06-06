// Earth: orbital period 365.25 Earth days, or 31557600 seconds
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years

class SpaceAge {
    private static final int SECONDS_IN_EARTH_YEAR = 31557600;
    private double seconds = 0;

    SpaceAge(double seconds) {
      this.seconds = seconds;
    }

    double onEarth() {
      return calculateAge(1);
    }

    double onMercury() {
      return calculateAge(0.2408467);
    }

    double onVenus() {
      return calculateAge(0.61519726);
    }

    double onMars() {
      return calculateAge(1.8808158);
    }

    double onJupiter() {
      return calculateAge(11.862615);
    }

    double onSaturn() {
      return calculateAge(29.47498);
    }

    double onUranus() {
      return calculateAge(84.016846);
    }

    double onNeptune() {
      return calculateAge(164.79132);
    }

    private double calculateAge(double conversionFactor) {
      return seconds / (SECONDS_IN_EARTH_YEAR * conversionFactor);
    }

}
