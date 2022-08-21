import 'package:english_words/utils/api/api.dart';
import 'package:english_words/utils/api/user/models.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

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
