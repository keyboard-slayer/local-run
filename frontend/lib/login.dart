import 'package:flutter/material.dart';
import 'package:moon_design/moon_design.dart';

class Login extends StatelessWidget {
  const Login({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
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
                MoonFormTextInput(
                  hintText: 'Username',
                  width: MediaQuery.of(context).size.width * 0.3,
                ),
                SizedBox(height: MediaQuery.of(context).size.height * 0.03),
                MoonFormTextInput(
                  hintText: 'Password',
                  obscureText: true,
                  width: MediaQuery.of(context).size.width * 0.3,
                ),
                SizedBox(height: MediaQuery.of(context).size.height * 0.04),
                MoonFilledButton(
                  buttonSize: MoonButtonSize.md,
                  onTap: () {},
                  label: const Text('Sign in'),
                  width: MediaQuery.of(context).size.width * 0.3,
                ),
                SizedBox(
                  width: MediaQuery.of(context).size.width * 0.3,
                  height: MediaQuery.of(context).size.height * 0.07,
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
    );
  }
}
