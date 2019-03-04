#! .venv/bin/python
# Simple client to connect a gRPC service to get information about the employee vacation.
import grpc
from protons.vacations_pb2 import VacationRequest
from protons.vacations_pb2_grpc import VacationServiceStub


def client():
    with grpc.insecure_channel("localhost:50051") as dial:
        client = VacationServiceStub(dial)
        response = client.WhenDoYouStartYourVacation(
            VacationRequest(name="Henrique Lopes"))
    print("Response: " + response.message)


if __name__ == "__main__":
    client()
