PYTHON=${shell python/.venv/bin/python}

procpython:
	${PYTHON} -m grpc_tools.protoc -I. --python_out=python --grpc_python_out=python protons/vacations.proto

procgo:
	protoc --go_out=plugins=grpc:. protons/vacations.proto
