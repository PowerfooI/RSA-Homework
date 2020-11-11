package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golangv/rsa"
	"net/http"
	"strconv"
)

func (s *Server) getKeys(c *gin.Context) {
	NumByteStr := c.Query("n_byte")
	NumByte, err := strconv.ParseInt(NumByteStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "wrong parameter",
		})
		return
	}
	public, private := rsa.GenRSAKeys(int(NumByte), 4)
	log.WithFields(log.Fields{
		"nByte":   NumByteStr,
		"public":  public,
		"private": private,
	}).Info("创建RSA密钥对成功")
	c.JSON(http.StatusOK, gin.H{
		"ok":         true,
		"publicKey":  genStrFromKeyType(public),
		"privateKey": genStrFromKeyType(private),
	})
}

func (s *Server) encodeMsg(c *gin.Context) {
	var err error
	body := struct {
		Key *Key   `json:"key"`
		Msg string `json:"msg"`
	}{}
	err = c.ShouldBind(&body)
	if err != nil {
		log.WithError(err).Error("加密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "wrong parameter",
		})
		return
	}
	k, err := genKeyTypeFromStr(body.Key)
	if err != nil {
		log.WithError(err).Error("加密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "bad pattern in keys",
		})
	}
	output := rsa.EncodeMsg(k, body.Msg)
	log.WithField("encodedMsg", output).Info("成功加密")
	c.JSON(http.StatusOK, gin.H{
		"ok":         true,
		"encodedMsg": output,
	})
}

func (s *Server) decodeMsg(c *gin.Context) {
	var err error
	body := struct {
		Key *Key   `json:"key"`
		Msg string `json:"msg"`
	}{}
	err = c.ShouldBind(&body)

	if err != nil {
		log.WithError(err).Error("解密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "wrong parameter",
		})
		return
	}
	k, err := genKeyTypeFromStr(body.Key)
	if err != nil {
		log.WithError(err).Error("解密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "bad pattern in keys",
		})
		return
	}
	output, err := rsa.DecodeMsg(k, body.Msg)
	if err != nil {
		log.WithError(err).Error("解密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "please use correct key for decoding",
		})
		return
	}
	log.WithField("decodedMsg", output).Info("成功解密")
	c.JSON(http.StatusOK, gin.H{
		"ok":         true,
		"decodedMsg": output,
	})
}

func (s *Server) signMsg(c *gin.Context) {
	var err error
	body := struct {
		Key *Key   `json:"key"`
		Msg string `json:"msg"`
	}{}
	err = c.ShouldBind(&body)

	if err != nil {
		log.WithError(err).Error("解密请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "wrong parameter",
		})
		return
	}
	k, err := genKeyTypeFromStr(body.Key)
	if err != nil {
		log.WithError(err).Error("签名请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "bad pattern in keys",
		})
		return
	}
	output := rsa.Sign(k, body.Msg)
	log.WithFields(log.Fields{
		"signature": output,
	}).Info("签名成功")
	c.JSON(http.StatusOK, gin.H{
		"ok":        true,
		"signature": output,
	})
}

func (s *Server) verifySignature(c *gin.Context) {
	var err error
	body := struct {
		Key       *Key   `json:"key"`
		Signature string `json:"signature"`
		Msg       string `json:"msg"`
	}{}
	err = c.ShouldBind(&body)

	if err != nil {
		log.WithError(err).Error("验证签名请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "wrong parameter",
		})
		return
	}
	k, err := genKeyTypeFromStr(body.Key)
	if err != nil {
		log.WithError(err).Error("验证签名请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "bad pattern in keys",
		})
		return
	}
	ok, err := rsa.Verify(k, body.Msg, body.Signature)
	if err != nil {
		log.WithError(err).Error("验证签名请求错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": "签名解密失败，签名被篡改或者密钥使用不正确",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":       true,
		"verified": ok,
	})
}
