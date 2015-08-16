# GoPush
## Introducton

GoPush is a message pushing library for Android devices and was inspired by [NoPush](https://github.com/SpongeBobSun/NoPush)

The Android client basicly is the same with NoPush and the server side is written by GoLang so we get a better concurrrent performance.

Also, GoPush is intent to create a free & open-source message pushing service.

GoPush can be deployed on your own server and work along with your own web applications.

## Project Status
Under developing and testing.

## Usage
## Server Interface
## Client Interface
Import the android side code to your project. Modify AndroidManifest.xml as below.

      <service android:name="sun.bob.nopush.NoPushService"
          android:enabled="true"
          android:process=":NoPushService"
          >
          <intent-filter>
              <category android:name="ANDROID.INTENT.CATEGORY.DEFAULT"/>
              <action android:name="sun.bob.nopush.pull_up_from_util"/>
              <action android:name="sun.bob.nopush.daemon_alive"/>
              <action android:name="sun.bob.nopush.message"/>
          </intent-filter>
          <intent-filter>
              <category android:name="android.intent.category.DEFAULT"/>
              <action android:name="sun.bob.nopush.daemon_pid"/>
          </intent-filter>
      </service>

After this, you can use below code to initialize the library.

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        NoPushUtil.initialize(getApplicationContext(),"your.server.ip.address",22333);
    }
## Design (How does it works?)
## Credits
[Jean.Jiang](https://github.com/JiangXuanYi)

[Bob.Sun](https://github.com/SpongeBobSun)
