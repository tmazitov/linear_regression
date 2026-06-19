NAME_TRAIN = train
NAME_PREDICT = predict

all: train predict

train:
	go build -o $(NAME_TRAIN) ./cmd/train

predict:
	go build -o $(NAME_PREDICT) ./cmd/predict

clean:
	go clean

fclean: clean
	rm -f $(NAME_TRAIN) $(NAME_PREDICT)

re: fclean all

.PHONY: all clean fclean re
