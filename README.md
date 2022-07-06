# GO-PAM-Demo

/etc/ssh/sshd_config
```
UsePAM yes
KbdInteractiveAuthentication yes
ChallengeResponseAuthentication yes
```

/etc/pam.d/sshd
```
auth required pam_demo.so
```
