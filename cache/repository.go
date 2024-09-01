package cache

import "context"

func GetAll(ctx context.Context, keys ...string) ([]interface{}, error) {
	ret := client.MGet(ctx, keys...)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return ret.Val(), nil
}

func SetAll(ctx context.Context, values ...interface{}) error {
	err := client.MSet(ctx, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

func RemoveAll(ctx context.Context, values ...string) error {
	err := client.Del(ctx, values...).Err()
	if err != nil {
		return err
	}
	return nil
}
