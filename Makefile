NAME_TRAIN = train
NAME_PREDICT = predict
NAME_PRECISION = precision

all: train predict precision

train:
	go build -o $(NAME_TRAIN) ./cmd/train

predict:
	go build -o $(NAME_PREDICT) ./cmd/predict

precision:
	go build -o $(NAME_PRECISION) ./cmd/precision

clean:
	go clean

fclean: clean
	rm -f $(NAME_TRAIN) $(NAME_PREDICT) $(NAME_PRECISION) weights.json

re: fclean all

.PHONY: all clean fclean re
