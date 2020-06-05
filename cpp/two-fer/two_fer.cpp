#include "two_fer.h"
#include <string>

using std::string;

namespace two_fer
{
  string two_fer(const string &name) {
    return "One for " + name + ", one for me.";
  }
} // namespace two_fer

