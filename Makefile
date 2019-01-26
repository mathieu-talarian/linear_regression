TRAINER=trainer
ESTIMATOR=estimator

E_SRC_F=main.go flags.go
T_SRC_F=main.go dataReader.go flags.go graph.go

E_DIR=src/estimator
T_DIR=src/trainer

E_SRC=$(addprefix $(E_DIR)/, $(E_SRC_F))
T_SRC=$(addprefix $(T_DIR)/, $(T_SRC_F))

all: $(TRAINER) $(ESTIMATOR)

$(TRAINER):
	go build -o $(TRAINER) $(T_SRC)

$(ESTIMATOR):
	go build -o $(ESTIMATOR) $(E_SRC)

fclean: 
	@rm -rfv $(TRAINER)
	@rm -rfv $(ESTIMATOR)

re: fclean all

trainer_debug: 
	go build -o debug_t -gcflags "-m -m -l" $(T_SRC)

estimator_debug:
	go build -o debug_e -gcflags "-m -m-l" $(E_SRC)