package mbox

import (
	"github.com/jbaikge/ingress-destroyed/message"
	"net/mail"
	"testing"
)

func TestSimple(t *testing.T) {
	msg := &message.Message{}
	if err := toMessage(simpleMboxMsg, msg); err != nil {
		t.Fatal(err)
	}
	t.Logf("ID:   %s", msg.Id)
	t.Logf("From: %s", msg.From)
	t.Logf("Date: %s", msg.Date)
}

func TestComplex(t *testing.T) {
	msg := &message.Message{}
	if err := toMessage(complexMboxMsg, msg); err != nil {
		t.Fatal(err)
	}
	t.Logf("ID:   %s", msg.Id)
	t.Logf("From: %s", msg.From)
	t.Logf("Date: %s", msg.Date)
}

func TestGetBoundary(t *testing.T) {
	exp := "089e013a2ef428047d04d381486c"
	h := mail.Header{
		"Content-Type": []string{"multipart/alternative; boundary=" + exp},
	}
	b, err := getBoundary(h)
	if err != nil {
		t.Fatal(err)
	}
	if b != exp {
		t.Fatalf("Mismatched boundary value. Got: %s", b)
	}
}

func TestGetContentType(t *testing.T) {
	exp := "multipart/alternative"
	h := mail.Header{
		"Content-Type": []string{exp + "; boundary=089e013a2ef428047d04d381486c"},
	}
	b, err := getContentType(h)
	if err != nil {
		t.Fatal(err)
	}
	if b != exp {
		t.Fatalf("Mismatched content-type. Got: %s", b)
	}
}

func TestComplexText(t *testing.T) {
	msg := &message.Message{}
	if err := toMessage(complexMboxMsg, msg); err != nil {
		t.Fatal(err)
	}
	t.Log(string(msg.Text))
}

func TestSimpleText(t *testing.T) {
	msg := &message.Message{}
	if err := toMessage(simpleMboxMsg, msg); err != nil {
		t.Fatal(err)
	}
	t.Log(string(msg.Text))
}

// Stolen from message_test in stdlib
var simpleMboxMsg = []byte(`From: John Doe <jdoe@machine.example>
To: Mary Smith <mary@example.net>
Subject: Saying Hello
Date: Fri, 21 Nov 1997 09:55:06 -0600
Message-ID: <1234@local.machine.example>

This is a message just to say hello.
So, "Hello".
`)

var complexMboxMsg = []byte(`From 3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com Thu May 23 22:00:04 2013
Return-path: <3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com>
Envelope-to: jake@localhost
Delivery-date: Thu, 23 May 2013 22:00:04 -0400
Received: from localhost ([127.0.0.1] helo=squeaker)
	by squeaker with esmtp (Exim 4.80)
	(envelope-from <3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com>)
	id 1UfhIZ-0002t0-PJ
	for jake@localhost; Thu, 23 May 2013 22:00:04 -0400
Delivered-To: bigjake2s@gmail.com
Received: from gmail-imap.l.google.com [173.194.76.108]
	by squeaker with IMAP (fetchmail-6.3.22)
	for <jake@localhost> (single-drop); Thu, 23 May 2013 22:00:03 -0400 (EDT)
Received: by 10.114.71.201 with SMTP id x9csp49093ldu; Thu, 17 Jan 2013
 11:48:42 -0800 (PST)
Received-SPF: pass (google.com: domain of
 3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com
 designates 10.58.188.106 as permitted sender) client-ip=10.58.188.106
Authentication-Results: mr.google.com; spf=pass (google.com: domain of
 3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com
 designates 10.58.188.106 as permitted sender)
 smtp.mail=3mVX4UAkJCbITWlSkhacWYeSad.UgeTaYbScWukYeSad.Uge@M3KW2WVRGUFZ5GODRSRYTGD7.apphosting.bounces.google.com;
 dkim=pass header.i=@google.com
X-Received: from mr.google.com ([10.58.188.106]) by 10.58.188.106 with SMTP id
 fz10mr3220952vec.37.1358452122115 (num_hops = 1); Thu, 17 Jan 2013 11:48:42
 -0800 (PST)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=google.com;
 s=20120113; h=mime-version:x-received:reply-to:x-google-appengine-app-id
 :x-google-appengine-app-id-alias:message-id:date:subject:from:to
 :content-type; bh=7ng+HUCa8UJ8NeLoz4G3W9CpBfpudjzzJ5rNIVndu10=;
 b=LaZREh/fmrOtJT+PGzXFGd9JJ5cEwcJdqduWt+rVn+wDwABLU9ESJR/xGK+YIuUMc8
 E4csvCq/9lx3CjQ9yoKqIINWOrSZ1BJouMHWTLf3ydkOr/dp7gxWHYR1VL49MaskKOSF
 YzSN7+h8UK/GKdWBUOl3iayWCQdtxZVWvRI9bFbqldXHEN1eOA5S0mZ6AZWMaBDlwP0e
 gO5qFAWI2K9Ik8ipDfcayg0NgPuefpVJE0VxiM2yqi+RUJNsfUMGmYmSThvIhq1wRdEY
 NZGSQ07UAJKMgD7+VKnAhuabAve9iPt1rgeoxSRi4rU3/CeSNw5binYvnaFECJx/P3YW yLrw==
MIME-Version: 1.0
X-Received: by 10.58.188.106 with SMTP id fz10mr2757662vec.37.1358452121665;
 Thu, 17 Jan 2013 11:48:41 -0800 (PST)
Reply-To: Niantic Project Operations <ingress-support@google.com>
X-Google-Appengine-App-Id: s~betaspike
X-Google-Appengine-App-Id-Alias: betaspike
Message-ID: <089e013a2ef42804b504d381486f@google.com>
Date: Thu, 17 Jan 2013 19:48:41 +0000
Subject: Ingress notification - Entities Destroyed by RedSoloCup
From: Niantic Project Operations <ingress-support@google.com>
To: jbaikge <bigjake2s@gmail.com>
Content-Type: multipart/alternative; boundary=089e013a2ef428047d04d381486c

--089e013a2ef428047d04d381486c
Content-Type: text/plain; charset=ISO-8859-1; format=flowed; delsp=yes

jbaikge,

1 Resonator(s) were destroyed by RedSoloCup at 11:48 hrs. - View location




Your Link has been destroyed by RedSoloCup at 11:48 hrs. - View start  
location - View end location



Your Link has been destroyed by RedSoloCup at 11:48 hrs. - View start  
location - View end location




------------------------------------------
Dashboard Contact

--089e013a2ef428047d04d381486c
Content-Type: text/html; charset=ISO-8859-1
Content-Transfer-Encoding: quoted-printable

jbaikge,<br/><br/>1 Resonator(s) were destroyed by RedSoloCup at 11:48 hrs.=
 - <a href=3D"http://www.ingress.com/intel?latE6=3D38807259&lngE6=3D-770639=
61&z=3D19">View location</a><br/> <br/><a href=3D"http://www.ingress.com/in=
tel?latE6=3D38807259&lngE6=3D-77063961&z=3D19"><img src=3D"http://lh4.ggpht=
.com/8g4RDnY0vnV9dTRen4dunE3wzaVRSA7vqOc24ToWMHF7JrHJ99rCWdRKubuquMDjrn8Lal=
GN8uns5p3bbM-m"/></a><br/><br/><br/>Your Link has been destroyed by RedSolo=
Cup at 11:48 hrs. - <a href=3D"http://www.ingress.com/intel?latE6=3D3880710=
9&lngE6=3D-77063212&z=3D19">View start location</a> - <a href=3D"http://www=
.ingress.com/intel?latE6=3D38807259&lngE6=3D-77063961&z=3D19">View end loca=
tion</a><br/><br/><a href=3D"http://www.ingress.com/intel?latE6=3D38807109&=
lngE6=3D-77063212&z=3D19"><img src=3D"http://www.panoramio.com/photos/small=
/59856083.jpg"/></a>&nbsp;<a href=3D"http://www.ingress.com/intel?latE6=3D3=
8807259&lngE6=3D-77063961&z=3D19"><img src=3D"http://lh4.ggpht.com/8g4RDnY0=
vnV9dTRen4dunE3wzaVRSA7vqOc24ToWMHF7JrHJ99rCWdRKubuquMDjrn8LalGN8uns5p3bbM-=
m"/></a><br/><br/>Your Link has been destroyed by RedSoloCup at 11:48 hrs. =
- <a href=3D"http://www.ingress.com/intel?latE6=3D38807259&lngE6=3D-7706396=
1&z=3D19">View start location</a> - <a href=3D"http://www.ingress.com/intel=
?latE6=3D38806683&lngE6=3D-77062351&z=3D19">View end location</a><br/><br/>=
<a href=3D"http://www.ingress.com/intel?latE6=3D38807259&lngE6=3D-77063961&=
z=3D19"><img src=3D"http://lh4.ggpht.com/8g4RDnY0vnV9dTRen4dunE3wzaVRSA7vqO=
c24ToWMHF7JrHJ99rCWdRKubuquMDjrn8LalGN8uns5p3bbM-m"/></a>&nbsp;<a href=3D"h=
ttp://www.ingress.com/intel?latE6=3D38806683&lngE6=3D-77062351&z=3D19"><img=
 src=3D"http://lh6.ggpht.com/Obr5YsHta6NbxhyGwyRirM9u8ejzcQqZS5SqTogQ_eQNZ0=
-9jq284ViGwtL8XuhsUnPlQTvxubhOALm25Wk"/></a><br/><br/><br/>----------------=
--------------------------<br/><a href=3D"http://www.ingress.com/intel">Das=
hboard</a>&nbsp;<a href=3D"http://support.google.com/ingress">Contact</a><b=
r/>
--089e013a2ef428047d04d381486c--
`)
