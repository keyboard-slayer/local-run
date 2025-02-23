import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:moon_design/moon_design.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import '../utils/dio.dart';

class Login extends StatefulWidget {
  const Login({super.key});

  @override
  State<Login> createState() => _LoginState();
}

class _LoginState extends State<Login> {
  bool _showMsg = false;
  String _msg = "";

  void setMsg(String newMsg) {
    setState(() {
      _msg = newMsg;
      _showMsg = true;
    });
  }

  void signIn(String username, String password) async {
    String? backend = dotenv.env['BACKEND_URL'];

    if (backend == null) {
      throw Exception('The backend is not set');
    }

    try {
      var response = await dio.post(
        '$backend/auth',
        data: {'username': username, 'password': password},
      );

      if (response.statusCode != 200) {
        setMsg("Unexpected backend error: ${response.statusMessage}");
      } else if (response.data['Status'] == "error") {
        setMsg(response.data['msg']);
      } else {
        setMsg('OK');
      }
    } on DioException catch (e) {
      if (e.type == DioExceptionType.connectionTimeout) {
        setMsg('Connection timed out!');
      } else if (e.type == DioExceptionType.badResponse) {
        setMsg('Received invalid status code: ${e.response?.statusCode}');
      } else {
        setMsg('Unexpected error: ${e.message}');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final formkey = GlobalKey<FormState>();
    final nameCtrl = TextEditingController();
    final pwdCtrl = TextEditingController();

    return Scaffold(
      body: Form(
        key: formkey,
        child: Stack(
          children: [
            Positioned(
              top: 0,
              right: 0,
              child: SizedBox(
                height: MediaQuery.of(context).size.height,
                width: MediaQuery.of(context).size.width * 0.558,
                child: Placeholder(),
              ),
            ),
            Positioned(
              top: MediaQuery.of(context).size.height * 0.27,
              left: MediaQuery.of(context).size.width * 0.08,
              child: Column(
                children: [
                  SizedBox(
                    width: MediaQuery.of(context).size.width * 0.3,
                    child: Visibility(
                      visible: _showMsg,
                      child: MoonAlert(
                        show: true,
                        color: context.moonColors!.chichi,
                        backgroundColor: context.moonColors!.chichi10,
                        leading: const Icon(
                          MoonIcons.notifications_alert_24_light,
                        ),
                        label: Text(_msg),
                        trailing: MoonButton.icon(
                          buttonSize: MoonButtonSize.xs,
                          onTap: () {
                            setState(() {
                              _showMsg = false;
                            });
                          },
                          icon: Icon(
                            MoonIcons.controls_close_small_24_light,
                            color: context.moonColors!.chichi,
                          ),
                        ),
                      ),
                    ),
                  ),
                  SizedBox(height: MediaQuery.of(context).size.height * 0.02),
                  MoonFormTextInput(
                    hintText: 'Username',
                    controller: nameCtrl,
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return 'Please enter your username';
                      }

                      return null;
                    },
                    width: MediaQuery.of(context).size.width * 0.3,
                  ),
                  SizedBox(height: MediaQuery.of(context).size.height * 0.03),
                  MoonFormTextInput(
                    hintText: 'Password',
                    controller: pwdCtrl,
                    obscureText: true,
                    width: MediaQuery.of(context).size.width * 0.3,
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return 'Please enter your password';
                      }

                      return null;
                    },
                  ),
                  SizedBox(height: MediaQuery.of(context).size.height * 0.04),

                  MoonFilledButton(
                    buttonSize: MoonButtonSize.md,
                    onTap:
                        () => {
                          if (formkey.currentState!.validate())
                            {signIn(nameCtrl.text, pwdCtrl.text)},
                        },
                    label: const Text('Sign in'),
                    width: MediaQuery.of(context).size.width * 0.3,
                  ),

                  SizedBox(
                    width: MediaQuery.of(context).size.width * 0.3,
                    height: 69,
                    child: Row(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: <Widget>[
                        const Expanded(child: Divider()),
                        Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 15),
                          child: Text(
                            "Or",
                            style: Theme.of(context).textTheme.labelMedium,
                          ),
                        ),
                        const Expanded(child: Divider()),
                      ],
                    ),
                  ),

                  MoonOutlinedButton(
                    buttonSize: MoonButtonSize.md,
                    onTap: () {},
                    label: const Text('Continue with passkey'),
                    width: MediaQuery.of(context).size.width * 0.3,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
