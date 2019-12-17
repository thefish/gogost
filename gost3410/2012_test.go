// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2020 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost3410

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"
	"testing/quick"
)

func TestGCL3Vectors(t *testing.T) {
	p := []byte{
		0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
		0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
		0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
		0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
		0xF1, 0xD8, 0x52, 0x74, 0x1A, 0xF4, 0x70, 0x4A,
		0x04, 0x58, 0x04, 0x7E, 0x80, 0xE4, 0x54, 0x6D,
		0x35, 0xB8, 0x33, 0x6F, 0xAC, 0x22, 0x4D, 0xD8,
		0x16, 0x64, 0xBB, 0xF5, 0x28, 0xBE, 0x63, 0x73,
	}
	q := []byte{
		0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
		0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
		0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
		0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
		0xA8, 0x2F, 0x2D, 0x7E, 0xCB, 0x1D, 0xBA, 0xC7,
		0x19, 0x90, 0x5C, 0x5E, 0xEC, 0xC4, 0x23, 0xF1,
		0xD8, 0x6E, 0x25, 0xED, 0xBE, 0x23, 0xC5, 0x95,
		0xD6, 0x44, 0xAA, 0xF1, 0x87, 0xE6, 0xE6, 0xDF,
	}
	a := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
	}
	b := []byte{
		0x1C, 0xFF, 0x08, 0x06, 0xA3, 0x11, 0x16, 0xDA,
		0x29, 0xD8, 0xCF, 0xA5, 0x4E, 0x57, 0xEB, 0x74,
		0x8B, 0xC5, 0xF3, 0x77, 0xE4, 0x94, 0x00, 0xFD,
		0xD7, 0x88, 0xB6, 0x49, 0xEC, 0xA1, 0xAC, 0x43,
		0x61, 0x83, 0x40, 0x13, 0xB2, 0xAD, 0x73, 0x22,
		0x48, 0x0A, 0x89, 0xCA, 0x58, 0xE0, 0xCF, 0x74,
		0xBC, 0x9E, 0x54, 0x0C, 0x2A, 0xDD, 0x68, 0x97,
		0xFA, 0xD0, 0xA3, 0x08, 0x4F, 0x30, 0x2A, 0xDC,
	}
	x := []byte{
		0x24, 0xD1, 0x9C, 0xC6, 0x45, 0x72, 0xEE, 0x30,
		0xF3, 0x96, 0xBF, 0x6E, 0xBB, 0xFD, 0x7A, 0x6C,
		0x52, 0x13, 0xB3, 0xB3, 0xD7, 0x05, 0x7C, 0xC8,
		0x25, 0xF9, 0x10, 0x93, 0xA6, 0x8C, 0xD7, 0x62,
		0xFD, 0x60, 0x61, 0x12, 0x62, 0xCD, 0x83, 0x8D,
		0xC6, 0xB6, 0x0A, 0xA7, 0xEE, 0xE8, 0x04, 0xE2,
		0x8B, 0xC8, 0x49, 0x97, 0x7F, 0xAC, 0x33, 0xB4,
		0xB5, 0x30, 0xF1, 0xB1, 0x20, 0x24, 0x8A, 0x9A,
	}
	y := []byte{
		0x2B, 0xB3, 0x12, 0xA4, 0x3B, 0xD2, 0xCE, 0x6E,
		0x0D, 0x02, 0x06, 0x13, 0xC8, 0x57, 0xAC, 0xDD,
		0xCF, 0xBF, 0x06, 0x1E, 0x91, 0xE5, 0xF2, 0xC3,
		0xF3, 0x24, 0x47, 0xC2, 0x59, 0xF3, 0x9B, 0x2C,
		0x83, 0xAB, 0x15, 0x6D, 0x77, 0xF1, 0x49, 0x6B,
		0xF7, 0xEB, 0x33, 0x51, 0xE1, 0xEE, 0x4E, 0x43,
		0xDC, 0x1A, 0x18, 0xB9, 0x1B, 0x24, 0x64, 0x0B,
		0x6D, 0xBB, 0x92, 0xCB, 0x1A, 0xDD, 0x37, 0x1E,
	}
	priv := []byte{
		0xD4, 0x8D, 0xA1, 0x1F, 0x82, 0x67, 0x29, 0xC6,
		0xDF, 0xAA, 0x18, 0xFD, 0x7B, 0x6B, 0x63, 0xA2,
		0x14, 0x27, 0x7E, 0x82, 0xD2, 0xDA, 0x22, 0x33,
		0x56, 0xA0, 0x00, 0x22, 0x3B, 0x12, 0xE8, 0x72,
		0x20, 0x10, 0x8B, 0x50, 0x8E, 0x50, 0xE7, 0x0E,
		0x70, 0x69, 0x46, 0x51, 0xE8, 0xA0, 0x91, 0x30,
		0xC9, 0xD7, 0x56, 0x77, 0xD4, 0x36, 0x09, 0xA4,
		0x1B, 0x24, 0xAE, 0xAD, 0x8A, 0x04, 0xA6, 0x0B,
	}
	pubX := []byte{
		0xE1, 0xEF, 0x30, 0xD5, 0x2C, 0x61, 0x33, 0xDD,
		0xD9, 0x9D, 0x1D, 0x5C, 0x41, 0x45, 0x5C, 0xF7,
		0xDF, 0x4D, 0x8B, 0x4C, 0x92, 0x5B, 0xBC, 0x69,
		0xAF, 0x14, 0x33, 0xD1, 0x56, 0x58, 0x51, 0x5A,
		0xDD, 0x21, 0x46, 0x85, 0x0C, 0x32, 0x5C, 0x5B,
		0x81, 0xC1, 0x33, 0xBE, 0x65, 0x5A, 0xA8, 0xC4,
		0xD4, 0x40, 0xE7, 0xB9, 0x8A, 0x8D, 0x59, 0x48,
		0x7B, 0x0C, 0x76, 0x96, 0xBC, 0xC5, 0x5D, 0x11,
	}
	pubY := []byte{
		0xEC, 0xBE, 0x77, 0x36, 0xA9, 0xEC, 0x35, 0x7F,
		0xF2, 0xFD, 0x39, 0x93, 0x1F, 0x4E, 0x11, 0x4C,
		0xB8, 0xCD, 0xA3, 0x59, 0x27, 0x0A, 0xC7, 0xF0,
		0xE7, 0xFF, 0x43, 0xD9, 0x41, 0x94, 0x19, 0xEA,
		0x61, 0xFD, 0x2A, 0xB7, 0x7F, 0x5D, 0x9F, 0x63,
		0x52, 0x3D, 0x3B, 0x50, 0xA0, 0x4F, 0x63, 0xE2,
		0xA0, 0xCF, 0x51, 0xB7, 0xC1, 0x3A, 0xDC, 0x21,
		0x56, 0x0F, 0x0B, 0xD4, 0x0C, 0xC9, 0xC7, 0x37,
	}
	digest := []byte{
		0x37, 0x54, 0xF3, 0xCF, 0xAC, 0xC9, 0xE0, 0x61,
		0x5C, 0x4F, 0x4A, 0x7C, 0x4D, 0x8D, 0xAB, 0x53,
		0x1B, 0x09, 0xB6, 0xF9, 0xC1, 0x70, 0xC5, 0x33,
		0xA7, 0x1D, 0x14, 0x70, 0x35, 0xB0, 0xC5, 0x91,
		0x71, 0x84, 0xEE, 0x53, 0x65, 0x93, 0xF4, 0x41,
		0x43, 0x39, 0x97, 0x6C, 0x64, 0x7C, 0x5D, 0x5A,
		0x40, 0x7A, 0xDE, 0xDB, 0x1D, 0x56, 0x0C, 0x4F,
		0xC6, 0x77, 0x7D, 0x29, 0x72, 0x07, 0x5B, 0x8C,
	}
	signature := []byte{
		0x10, 0x81, 0xB3, 0x94, 0x69, 0x6F, 0xFE, 0x8E,
		0x65, 0x85, 0xE7, 0xA9, 0x36, 0x2D, 0x26, 0xB6,
		0x32, 0x5F, 0x56, 0x77, 0x8A, 0xAD, 0xBC, 0x08,
		0x1C, 0x0B, 0xFB, 0xE9, 0x33, 0xD5, 0x2F, 0xF5,
		0x82, 0x3C, 0xE2, 0x88, 0xE8, 0xC4, 0xF3, 0x62,
		0x52, 0x60, 0x80, 0xDF, 0x7F, 0x70, 0xCE, 0x40,
		0x6A, 0x6E, 0xEB, 0x1F, 0x56, 0x91, 0x9C, 0xB9,
		0x2A, 0x98, 0x53, 0xBD, 0xE7, 0x3E, 0x5B, 0x4A,
		0x2F, 0x86, 0xFA, 0x60, 0xA0, 0x81, 0x09, 0x1A,
		0x23, 0xDD, 0x79, 0x5E, 0x1E, 0x3C, 0x68, 0x9E,
		0xE5, 0x12, 0xA3, 0xC8, 0x2E, 0xE0, 0xDC, 0xC2,
		0x64, 0x3C, 0x78, 0xEE, 0xA8, 0xFC, 0xAC, 0xD3,
		0x54, 0x92, 0x55, 0x84, 0x86, 0xB2, 0x0F, 0x1C,
		0x9E, 0xC1, 0x97, 0xC9, 0x06, 0x99, 0x85, 0x02,
		0x60, 0xC9, 0x3B, 0xCB, 0xCD, 0x9C, 0x5C, 0x33,
		0x17, 0xE1, 0x93, 0x44, 0xE1, 0x73, 0xAE, 0x36,
	}
	c, err := NewCurve(
		bytes2big(p),
		bytes2big(q),
		bytes2big(a),
		bytes2big(b),
		bytes2big(x),
		bytes2big(y),
		nil,
		nil,
	)
	if err != nil {
		t.FailNow()
	}
	prv, err := NewPrivateKey(c, Mode2012, priv)
	if err != nil {
		t.FailNow()
	}
	pub, err := prv.PublicKey()
	if err != nil {
		t.FailNow()
	}
	if bytes.Compare(pub.Raw()[:64], pubX) != 0 {
		t.FailNow()
	}
	if bytes.Compare(pub.Raw()[64:], pubY) != 0 {
		t.FailNow()
	}
	ourSign, err := prv.SignDigest(digest, rand.Reader)
	if err != nil {
		t.FailNow()
	}
	valid, err := pub.VerifyDigest(digest, ourSign)
	if err != nil || !valid {
		t.FailNow()
	}
	valid, err = pub.VerifyDigest(digest, signature)
	if err != nil || !valid {
		t.FailNow()
	}
}

func TestRandom2012(t *testing.T) {
	c := CurveIdtc26gost341012512paramSetA()
	f := func(prvRaw [64 - 1]byte, digest [64]byte) bool {
		prv, err := NewPrivateKey(
			c,
			Mode2012,
			append([]byte{0xde}, prvRaw[:]...),
		)
		if err != nil {
			return false
		}
		pub, err := prv.PublicKey()
		if err != nil {
			return false
		}
		pubRaw := pub.Raw()
		pub, err = NewPublicKey(c, Mode2012, pubRaw)
		if err != nil {
			return false
		}
		sign, err := prv.SignDigest(digest[:], rand.Reader)
		if err != nil {
			return false
		}
		valid, err := pub.VerifyDigest(digest[:], sign)
		if err != nil {
			return false
		}
		return valid
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Test vectors from RFC 8133
func TestSESPAKE(t *testing.T) {
	type vector struct {
		curve          *Curve
		qIndX          string
		qIndY          string
		fPW            string
		xExpected      string
		yExpected      string
		alpha          string
		xAlphaExpected string
		yAlphaExpected string
		beta           string
		xBetaExpected  string
		yBetaExpected  string
	}
	vectors := []vector{
		vector{
			curve:          CurveIdGostR34102001CryptoProAParamSet(),
			qIndX:          "A69D51CAF1A309FA9E9B66187759B0174C274E080356F23CFCBFE84D396AD7BB",
			qIndY:          "5D26F29ECC2E9AC0404DCF7986FA55FE94986362170F54B9616426A659786DAC",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F18127067",
			xExpected:      "59495655D1E7C7424C622485F575CCF121F3122D274101E8AB734CC9C9A9B45E",
			yExpected:      "48D1C311D33C9B701F3B03618562A4A07A044E3AF31E3999E67B487778B53C62",
			alpha:          "1F2538097D5A031FA68BBB43C84D12B3DE47B7061C0D5E24993E0C873CDBA6B3",
			xAlphaExpected: "BBC77CF42DC1E62D06227935379B4AA4D14FEA4F565DDF4CB4FA4D31579F9676",
			yAlphaExpected: "8E16604A4AFDF28246684D4996274781F6CB80ABBBA1414C1513EC988509DABF",
			beta:           "DC497D9EF6324912FD367840EE509A2032AEDB1C0A890D133B45F596FCCBD45D",
			xBetaExpected:  "6097341C1BE388E83E7CA2DF47FAB86E2271FD942E5B7B2EB2409E49F742BC29",
			yBetaExpected:  "C81AA48BDB4CA6FA0EF18B9788AE25FE30857AA681B3942217F9FED151BAB7D0",
		},
		vector{
			curve:          CurveIdGostR34102001CryptoProBParamSet(),
			qIndX:          "3D715A874A4B17CB3B517893A9794A2B36C89D2FFC693F01EE4CC27E7F49E399",
			qIndY:          "1C5A641FCF7CE7E87CDF8CEA38F3DB3096EACE2FAD158384B53953365F4FE7FE",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F18127067",
			xExpected:      "6DC2AE26BC691FCA5A73D9C452790D15E34BA5404D92955B914C8D2662ABB985",
			yExpected:      "3B02AAA9DD65AE30C335CED12F3154BBAC059F66B088306747453EDF6E5DB077",
			alpha:          "499D72B90299CAB0DA1F8BE19D9122F622A13B32B730C46BD0664044F2144FAD",
			xAlphaExpected: "61D6F916DB717222D74877F179F7EBEF7CD4D24D8C1F523C048E34A1DF30F8DD",
			yAlphaExpected: "3EC48863049CFCFE662904082E78503F4973A4E105E2F1B18C69A5E7FB209000",
			beta:           "0F69FF614957EF83668EDC2D7ED614BE76F7B253DB23C5CC9C52BF7DF8F4669D",
			xBetaExpected:  "33BC6F7E9C0BA10CFB2B72546C327171295508EA97F8C8BA9F890F2478AB4D6C",
			yBetaExpected:  "75D57B396C396F492F057E9222CCC686437A2AAD464E452EF426FC8EEED1A4A6",
		},
		vector{
			curve:          CurveIdGostR34102001CryptoProCParamSet(),
			qIndX:          "1E36383E43BB6CFA2917167D71B7B5DD3D6D462B43D7C64282AE67DFBEC2559D",
			qIndY:          "137478A9F721C73932EA06B45CF72E37EB78A63F29A542E563C614650C8B6399",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F18127067",
			xExpected:      "945821DAF91E158B839939630655A3B21FF3E146D27041E86C05650EB3B46B59",
			yExpected:      "3A0C2816AC97421FA0E879605F17F0C9C3EB734CFF196937F6284438D70BDC48",
			alpha:          "3A54AC3F19AD9D0B1EAC8ACDCEA70E581F1DAC33D13FEAFD81E762378639C1A8",
			xAlphaExpected: "96B7F09C94D297C257A7DA48364C0076E59E48D221CBA604AE111CA3933B446A",
			yAlphaExpected: "54E4953D86B77ECCEB578500931E822300F7E091F79592CA202A020D762C34A6",
			beta:           "448781782BF7C0E52A1DD9E6758FD3482D90D3CFCCF42232CF357E59A4D49FD4",
			xBetaExpected:  "4B9C0AB55A938121F282F48A2CC4396EB16E7E0068B495B0C1DD4667786A3EB7",
			yBetaExpected:  "223460AA8E09383E9DF9844C5A0F2766484738E5B30128A171B69A77D9509B96",
		},
		vector{
			curve:          CurveIdtc26gost341012512paramSetA(),
			qIndX:          "2A17F8833A32795327478871B5C5E88AEFB91126C64B4B8327289BEA62559425D18198F133F400874328B220C74497CD240586CB249E158532CB8090776CD61C",
			qIndY:          "728F0C4A73B48DA41CE928358FAD26B47A6E094E9362BAE82559F83CDDC4EC3A4676BD3707EDEAF4CD85E99695C64C241EDC622BE87DC0CF87F51F4367F723C5",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F181270671C6213E3930EFDDA26451792C6208122EE60D200520D695DFD9F5F0FD5ABA702",
			xExpected:      "0C0AB53D0E0A9C607CAD758F558915A0A7DC5DC87B45E9A58FDDF30EC3385960283E030CD322D9E46B070637785FD49D2CD711F46807A24C40AF9A42C8E2D740",
			yExpected:      "DF93A8012B86D3A3D4F8A4D487DA15FC739EB31B20B3B0E8C8C032AAF8072C6337CF7D5B404719E5B4407C41D9A3216A08CA69C271484E9ED72B8AAA52E28B8B",
			alpha:          "3CE54325DB52FE798824AEAD11BB16FA766857D04A4AF7D468672F16D90E7396046A46F815693E85B1CE5464DA9270181F82333B0715057BBE8D61D400505F0E",
			xAlphaExpected: "B93093EB0FCC463239B7DF276E09E592FCFC9B635504EA4531655D76A0A3078E2B4E51CFE2FA400CC5DE9FBE369DB204B3E8ED7EDD85EE5CCA654C1AED70E396",
			yAlphaExpected: "809770B8D910EA30BD2FA89736E91DC31815D2D9B31128077EEDC371E9F69466F497DC64DD5B1FADC587F860EE256109138C4A9CD96B628E65A8F590520FC882",
			beta:           "B5C286A79AA8E97EC0E19BC1959A1D15F12F8C97870BA9D68CC12811A56A3BB11440610825796A49D468CDC9C2D02D76598A27973D5960C5F50BCE28D8D345F4",
			xBetaExpected:  "238B38644E440452A99FA6B93D9FD7DA0CB83C32D3C1E3CFE5DF5C3EB0F9DB91E588DAEDC849EA2FB867AE855A21B4077353C0794716A6480995113D8C20C7AF",
			yBetaExpected:  "B2273D5734C1897F8D15A7008B862938C8C74CA7E877423D95243EB7EBD02FD2C456CF9FC956F078A59AA86F19DD1075E5167E4ED35208718EA93161C530ED14",
		},
		vector{
			curve:          CurveIdtc26gost341012512paramSetB(),
			qIndX:          "7E1FAE8285E035BEC244BEF2D0E5EBF436633CF50E55231DEA9C9CF21D4C8C33DF85D4305DE92971F0A4B4C07E00D87BDBC720EB66E49079285AAF12E0171149",
			qIndY:          "2CC89998B875D4463805BA0D858A196592DB20AB161558FF2F4EF7A85725D20953967AE621AFDEAE89BB77C83A2528EF6FCE02F68BDA4679D7F2704947DBC408",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F181270671C6213E3930EFDDA26451792C6208122EE60D200520D695DFD9F5F0FD5ABA702",
			xExpected:      "7D03E65B8050D1E12CBB601A17B9273B0E728F5021CD47C8A4DD822E4627BA5F9C696286A2CDDA9A065509866B4DEDEDC4A118409604AD549F87A60AFA621161",
			yExpected:      "16037DAD45421EC50B00D50BDC6AC3B85348BC1D3A2F85DB27C3373580FEF87C2C743B7ED30F22BE22958044E716F93A61CA3213A361A2797A16A3AE62957377",
			alpha:          "715E893FA639BF341296E0623E6D29DADF26B163C278767A7982A989462A3863FE12AEF8BD403D59C4DC4720570D4163DB0805C7C10C4E818F9CB785B04B9997",
			xAlphaExpected: "10C479EA1C04D3C2C02B0576A9C42D96226FF033C1191436777F66916030D87D02FB93738ED7669D07619FFCE7C1F3C4DB5E5DF49E2186D6FA1E2EB5767602B9",
			yAlphaExpected: "039F6044191404E707F26D59D979136A831CCE43E1C5F0600D1DDF8F39D0CA3D52FBD943BF04DDCED1AA2CE8F5EBD7487ACDEF239C07D015084D796784F35436",
			beta:           "30FA8C2B4146C2DBBE82BED04D7378877E8C06753BD0A0FF71EBF2BEFE8DA8F3DC0836468E2CE7C5C961281B6505140F8407413F03C2CB1D201EA1286CE30E6D",
			xBetaExpected:  "34C0149E7BB91AE377B02573FCC48AF7BFB7B16DEB8F9CE870F384688E3241A3A868588CC0EF4364CCA67D17E3260CD82485C202ADC76F895D5DF673B1788E67",
			yBetaExpected:  "608E944929BD643569ED5189DB871453F13333A1EAF82B2FE1BE8100E775F13DD9925BD317B63BFAF05024D4A738852332B64501195C1B2EF789E34F23DDAFC5",
		},
		vector{
			curve:          CurveIdtc26gost34102012256paramSetA(),
			qIndX:          "B51ADF93A40AB15792164FAD3352F95B66369EB2A4EF5EFAE32829320363350E",
			qIndY:          "74A358CC08593612F5955D249C96AFB7E8B0BB6D8BD2BBE491046650D822BE18",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F18127067",
			xExpected:      "DBF99827078956812FA48C6E695DF589DEF1D18A2D4D35A96D75BF6854237629",
			yExpected:      "9FDDD48BFBC57BEE1DA0CFF282884F284D471B388893C48F5ECB02FC18D67589",
			alpha:          "147B72F6684FB8FD1B418A899F7DBECAF5FCE60B13685BAA95328654A7F0707F",
			xAlphaExpected: "33FBAC14EAE538275A769417829C431BD9FA622B6F02427EF55BD60EE6BC2888",
			yAlphaExpected: "22F2EBCF960A82E6CDB4042D3DDDA511B2FBA925383C2273D952EA2D406EAE46",
			beta:           "30D5CFADAA0E31B405E6734C03EC4C5DF0F02F4BA25C9A3B320EE6453567B4CB",
			xBetaExpected:  "2B2D89FAB735433970564F2F28CFA1B57D640CB902BC6334A538F44155022CB2",
			yBetaExpected:  "10EF6A82EEF1E70F942AA81D6B4CE5DEC0DDB9447512962874870E6F2849A96F",
		},
		vector{
			curve:          CurveIdtc26gost34102012512paramSetC(),
			qIndX:          "489C91784E02E98F19A803ABCA319917F37689E5A18965251CE2FF4E8D8B298F5BA7470F9E0E713487F96F4A8397B3D09A270C9D367EB5E0E6561ADEEB51581D",
			qIndY:          "684EA885ACA64EAF1B3FEE36C0852A3BE3BD8011B0EF18E203FF87028D6EB5DB2C144A0DCC71276542BFD72CA2A43FA4F4939DA66D9A60793C704A8C94E16F18",
			fPW:            "BD04673F7149B18E98155BD1E2724E71D0099AA25174F792D3326C6F181270671C6213E3930EFDDA26451792C6208122EE60D200520D695DFD9F5F0FD5ABA702",
			xExpected:      "0185AE6271A81BB7F236A955F7CAA26FB63849813C0287D96C83A15AE6B6A86467AB13B6D88CE8CD7DC2E5B97FF5F28FAC2C108F2A3CF3DB5515C9E6D7D210E8",
			yExpected:      "ED0220F92EF771A71C64ECC77986DB7C03D37B3E2AB3E83F32CE5E074A762EC08253C9E2102B87532661275C4B1D16D2789CDABC58ACFDF7318DE70AB64F09B8",
			alpha:          "332F930421D14CFE260042159F18E49FD5A54167E94108AD80B1DE60B13DE7999A34D611E63F3F870E5110247DF8EC7466E648ACF385E52CCB889ABF491EDFF0",
			xAlphaExpected: "561655966D52952E805574F4281F1ED3A2D498932B00CBA9DECB42837F09835BFFBFE2D84D6B6B242FE7B57F92E1A6F2413E12DDD6383E4437E13D72693469AD",
			yAlphaExpected: "F6B18328B2715BD7F4178615273A36135BC0BF62F7D8BB9F080164AD36470AD03660F51806C64C6691BADEF30F793720F8E3FEAED631D6A54A4C372DCBF80E82",
			beta:           "38481771E7D054F96212686B613881880BD8A6C89DDBC656178F014D2C093432A033EE10415F13A160D44C2AD61E6E2E05A7F7EC286BCEA3EA4D4D53F8634FA2",
			xBetaExpected:  "B7C5818687083433BC1AFF61CB5CA79E38232025E0C1F123B8651E62173CE6873F3E6FFE7281C2E45F4F524F66B0C263616ED08FD210AC4355CA3292B51D71C3",
			yBetaExpected:  "497F14205DBDC89BDDAF50520ED3B1429AD30777310186BE5E68070F016A44E0C766DB08E8AC23FBDFDE6D675AA4DF591EB18BA0D348DF7AA40973A2F1DCFA55",
		},
	}
	for _, vector := range vectors {
		raw, _ := hex.DecodeString(vector.qIndX)
		qIndX := bytes2big(raw)
		raw, _ = hex.DecodeString(vector.qIndY)
		qIndY := bytes2big(raw)
		raw, _ = hex.DecodeString(vector.fPW)
		reverse(raw)
		f := bytes2big(raw)
		x, y, err := vector.curve.Exp(f, qIndX, qIndY)
		if err != nil {
			t.FailNow()
		}
		raw, _ = hex.DecodeString(vector.xExpected)
		if bytes.Compare(x.Bytes(), raw) != 0 {
			t.FailNow()
		}
		raw, _ = hex.DecodeString(vector.yExpected)
		if bytes.Compare(y.Bytes(), raw) != 0 {
			t.FailNow()
		}

		raw, _ = hex.DecodeString(vector.alpha)
		alpha := bytes2big(raw)
		x, y, err = vector.curve.Exp(alpha, vector.curve.X, vector.curve.Y)
		raw, _ = hex.DecodeString(vector.xAlphaExpected)
		if bytes.Compare(x.Bytes(), raw) != 0 {
			t.FailNow()
		}
		raw, _ = hex.DecodeString(vector.yAlphaExpected)
		if bytes.Compare(y.Bytes(), raw) != 0 {
			t.FailNow()
		}

		raw, _ = hex.DecodeString(vector.beta)
		beta := bytes2big(raw)
		x, y, err = vector.curve.Exp(beta, vector.curve.X, vector.curve.Y)
		raw, _ = hex.DecodeString(vector.xBetaExpected)
		if bytes.Compare(x.Bytes(), raw) != 0 {
			t.FailNow()
		}
		raw, _ = hex.DecodeString(vector.yBetaExpected)
		if bytes.Compare(y.Bytes(), raw) != 0 {
			t.FailNow()
		}
	}
}

// Twisted Edwards to Weierstrass coordinates conversion and vice versa
func TestUVXYConversion(t *testing.T) {
	f := func(c *Curve, uRaw, vRaw []byte) {
		u := bytes2big(uRaw)
		v := bytes2big(vRaw)
		x, y := UV2XY(c, u, v)
		if x.Cmp(c.X) != 0 {
			t.FailNow()
		}
		if y.Cmp(c.Y) != 0 {
			t.FailNow()
		}
		uGot, vGot := XY2UV(c, c.X, c.Y)
		if u.Cmp(uGot) != 0 {
			t.FailNow()
		}
		if v.Cmp(vGot) != 0 {
			t.FailNow()
		}
	}
	raw, _ := hex.DecodeString("60CA1E32AA475B348488C38FAB07649CE7EF8DBE87F22E81F92B2592DBA300E7")
	f(CurveIdtc26gost34102012256paramSetA(), []byte{0x0D}, raw)
	raw, _ = hex.DecodeString("469AF79D1FB1F5E16B99592B77A01E2A0FDFB0D01794368D9A56117F7B38669522DD4B650CF789EEBF068C5D139732F0905622C04B2BAAE7600303EE73001A3D")
	f(CurveIdtc26gost34102012512paramSetC(), []byte{0x12}, raw)
}

func BenchmarkSign2012(b *testing.B) {
	c := CurveIdtc26gost341012512paramSetA()
	prv, err := GenPrivateKey(c, Mode2012, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	digest := make([]byte, 64)
	rand.Read(digest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prv.SignDigest(digest, rand.Reader)
	}
}

func BenchmarkVerify2012(b *testing.B) {
	c := CurveIdtc26gost341012512paramSetA()
	prv, err := GenPrivateKey(c, Mode2012, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	digest := make([]byte, 64)
	rand.Read(digest)
	sign, err := prv.SignDigest(digest, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	pub, err := prv.PublicKey()
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.VerifyDigest(digest, sign)
	}
}
