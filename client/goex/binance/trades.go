package binance

//type ListSubscriptionsService struct {
//	params []interface{}
//	c      *goex.Client
//}
//
//func (sub *ListSubscriptionsService) Do(ctx context.Context, ID uint) error {
//	msg := &RpcReq{
//		Method: "LIST_SUBSCRIPTIONS",
//		Params: sub.params,
//		ID:     ID,
//	}
//	data, err := json.Marshal(msg)
//	if err != nil {
//		return err
//	}
//	sub.c.Send(data)
//	return nil
//}
//
//type SetPropertyService struct {
//	params []interface{}
//	c      *goex.Client
//}
//
//func (s SetPropertyService) Combined(v bool) {
//	s.params = append(s.params, "combined")
//	s.params = append(s.params, v)
//}
//
//func (s *SetPropertyService) Do(ctx context.Context, ID uint) error {
//	msg := &RpcReq{
//		Method: "SET_PROPERTY",
//		Params: s.params,
//		ID:     ID,
//	}
//	data, err := json.Marshal(msg)
//	if err != nil {
//		return err
//	}
//	s.c.Send(data)
//	return nil
//}
//
//type GetPropertyService struct {
//	params []interface{}
//	c      *goex.Client
//}
//
//func (s GetPropertyService) Combined(v bool) {
//	s.params = append(s.params, "combined")
//	s.params = append(s.params, v)
//}
//
//func (s *GetPropertyService) Do(ctx context.Context, ID uint) error {
//	msg := &RpcReq{
//		Method: "GET_PROPERTY",
//		Params: s.params,
//		ID:     ID,
//	}
//	data, err := json.Marshal(msg)
//	if err != nil {
//		return err
//	}
//	s.c.Send(data)
//	return nil
//}
