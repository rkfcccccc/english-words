import 'package:english_words/utils/api/api.dart';
import 'package:english_words/utils/api/user/models.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

const storage = FlutterSecureStorage();

Future<SignupResponse> signup(
  String email,
  password, {
  VerificationData? verification,
}) async {
  final data = await request("post", "/user/signup", body: {
    "email": email,
    "password": password,
    "verification": verification?.toMap()
  });

  return SignupResponse.fromMap(data);
}

Future<CredentialsResponse> login(String email, password) async {
  final data = await request("post", "/user/login", body: {
    "email": email,
    "password": password,
  });

  return CredentialsResponse.fromMap(data);
}

var refreshLock = false;
Future<void> refresh() async {
  if (refreshLock) {
    while (refreshLock) {
      await Future.delayed(const Duration(milliseconds: 100));
    }

    return;
  }

  final data = await request("post", "/user/refresh", body: {
    "token": await storage.read(key: "refresh"),
  });

  final credentials = CredentialsResponse.fromMap(data);
  await storeCredentials(credentials.jwt, credentials.refresh);
}

Future<String?> getAccessToken() async {
  if (refreshLock) {
    while (refreshLock) {
      await Future.delayed(const Duration(milliseconds: 100));
    }
  }

  final token = await storage.read(key: "jwt");
  return token;
}

Future<void> storeCredentials(String jwt, String refresh) async {
  await Future.wait([
    storage.write(key: "jwt", value: jwt),
    storage.write(key: "refresh", value: refresh),
  ]);
}
