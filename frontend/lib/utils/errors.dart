enum FailureError {
  unknownError,
  serverInternalError,

  noInternet,

  invalidEmail,
  invalidPassword,

  tooManyRequests,
  noAttemptsLeft,

  alreadyExist,
  notFound,

  unauthorized,
  wrongCode,
}

class AppError {
  final FailureError key;
  final String? message;

  const AppError({required this.key, this.message});

  @override
  String toString() => message ?? key.toString();
}

void throwSomeAppError(String? errorName) {
  switch (errorName) {
    case "INVALID_EMAIL":
      throw const AppError(
        key: FailureError.invalidEmail,
        message: "Invalid email",
      );
    case "INVALID_PASSWORD":
      throw const AppError(
        key: FailureError.invalidPassword,
        message: "Invalid password",
      );
    case "UNAUTHORIZED":
      throw const AppError(
        key: FailureError.unauthorized,
        message: "Invalid email or password",
      );
    case "ALREADY_EXISTS":
      throw const AppError(
        key: FailureError.notFound,
        message: "Email is already in use",
      );
    case "NOT_FOUND":
      throw const AppError(
        key: FailureError.notFound,
        message: "Not found",
      );
    case "NO_ATTEMPTS_LEFT":
      throw const AppError(
        key: FailureError.noAttemptsLeft,
        message: "No attempts left, try resending the code.",
      );
    case "WRONG_CODE":
      throw const AppError(
        key: FailureError.wrongCode,
        message: "Wrong code",
      );
    default:
      print("unhanlded error name $errorName");
      throw const AppError(
        key: FailureError.unknownError,
        message: "Unknown error",
      );
  }
}
