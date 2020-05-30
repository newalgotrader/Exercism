# Earth: orbital period 365.25 Earth days, or 31557600 seconds
# Mercury: orbital period 0.2408467 Earth years
# Venus: orbital period 0.61519726 Earth years
# Mars: orbital period 1.8808158 Earth years
# Jupiter: orbital period 11.862615 Earth years
# Saturn: orbital period 29.447498 Earth years
# Uranus: orbital period 84.016846 Earth years
# Neptune: orbital period 164.79132 Earth years

PRECISION = 2
SECONDS_IN_EARTH_YEAR = 31557600

class SpaceAge:
    def __init__(self, seconds):
        self.age = seconds

    def on_earth(self):
        return self.__calculate_age(self.age, 1)

    def on_mercury(self):
        return self.__calculate_age(self.age, 0.2408467)

    def on_venus(self):
        return self.__calculate_age(self.age, 0.61519726)

    def on_mars(self):
        return self.__calculate_age(self.age, 1.8808158)

    def on_jupiter(self):
        return self.__calculate_age(self.age, 11.862615)

    def on_saturn(self):
        return self.__calculate_age(self.age, 29.47498)

    def on_uranus(self):
        return self.__calculate_age(self.age, 84.016846)

    def on_neptune(self):
        return self.__calculate_age(self.age, 164.79132)

    def __calculate_age(_, age, conversion_factor):
        return round(age / (SECONDS_IN_EARTH_YEAR * conversion_factor),
                PRECISION)
