import 'dart:convert';

class VerificationData {
  final String requestId;
  final int code;

  VerificationData({
    required this.requestId,
    required this.code,
  });

  Map<String, dynamic> toMap() {
    return {
      'request_id': requestId,
      'code': code,
    };
  }

  factory VerificationData.fromMap(Map<String, dynamic> map) {
    return VerificationData(
      requestId: map['request_id'] ?? '',
      code: map['code']?.toInt() ?? 0,
    );
  }

  factory VerificationData.fromJson(String source) =>
      VerificationData.fromMap(json.decode(source));

  @override
  String toString() => 'VerificationData(requestId: $requestId, code: $code)';
}

class CredentialsResponse {
  final String jwt;
  final String refresh;

  CredentialsResponse({
    required this.jwt,
    required this.refresh,
  });

  factory CredentialsResponse.fromMap(Map<String, dynamic> map) {
    return CredentialsResponse(
      jwt: map['jwt'] ?? '',
      refresh: map['refresh'] ?? '',
    );
  }

  factory CredentialsResponse.fromJson(String source) =>
      CredentialsResponse.fromMap(json.decode(source));

  @override
  String toString() => 'CredentialsResponse(jwt: $jwt, refresh: $refresh)';
}

class SignupResponse {
  final String? jwt;
  final String? refresh;
  final String? requestId;
  SignupResponse({
    this.jwt,
    this.refresh,
    this.requestId,
  });

  factory SignupResponse.fromMap(Map<String, dynamic> map) {
    return SignupResponse(
      jwt: map['jwt'],
      refresh: map['refresh'],
      requestId: map['request_id'],
    );
  }

  @override
  String toString() =>
      'SignupResponse(jwt: $jwt, refresh: $refresh, requestId: $requestId)';
}
