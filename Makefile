PYTHON=${shell python/.venv/bin/python}
PIP=${shell python/.venv/bin/pip}

procpython:
	${PYTHON} -m grpc_tools.protoc -I. --python_out=python --grpc_python_out=python protons/vacations.proto

procgo:
	protoc --go_out=plugins=grpc:go protons/vacations.proto

setup:
	virtualenv --python=python3 .venv
	${PIP} install -r python/requirements.txt
