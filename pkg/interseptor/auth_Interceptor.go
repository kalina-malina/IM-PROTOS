package interceptor

import (
	"context"
	"fmt"
	"strconv"

	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthData struct {
	UserID      uint64
	Role        string
	AccessToken string
	SessionID   string
	RequestID   string
}

type ctxKey string

const ctxKeyAuth ctxKey = "auth_data"

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "ошибка в interceptor: %v", r)
		}
	}()

	requestID := uuid.New().String()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	accessTokens := md.Get("authorization")

	var accessToken string
	if len(accessTokens) > 0 {
		accessToken = strings.TrimPrefix(accessTokens[0], "Bearer ")
	}

	userIDs := md.Get("x-user-id")
	var userID uint64
	if len(userIDs) > 0 {
		userID, err = strconv.ParseUint(userIDs[0], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "неверный user id: %v", err)
		}
	}

	roles := md.Get("x-role")
	var role string
	if len(roles) > 0 {
		role = roles[0]
	}

	sessionIDs := md.Get("x-session-id")
	var sessionID string
	if len(sessionIDs) > 0 {
		sessionID = sessionIDs[0]
	}

	auth := &AuthData{
		UserID:      userID,
		Role:        role,
		AccessToken: accessToken,
		SessionID:   sessionID,
		RequestID:   requestID,
	}
	ctx = context.WithValue(ctx, ctxKeyAuth, auth)

	return handler(ctx, req)
}

func GetAuthData(ctx context.Context) *AuthData {
	val := ctx.Value(ctxKeyAuth)

	if val == nil {
		return &AuthData{}
	}

	if auth, ok := val.(*AuthData); ok {
		return auth
	}
	auth := val.(*AuthData)
	auth.UserID = uint64(auth.UserID)

	return &AuthData{}
}
