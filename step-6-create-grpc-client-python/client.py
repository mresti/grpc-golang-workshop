import grpc

import books_pb2 as pb
import books_pb2_grpc

def list_books(stub):
    # Exception handling.
    try:
        # List books
        books = stub.List(pb.Empty)
        print(books)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('ListUsers failed with {0}: {1}'.format(e.code(), e.details()))

def get_book(stub):
    # Exception handling.
    try:
        # List books
        response = stub.Get(pb.BookIdRequest(id=123))
        print(response)

        response = stub.Get(pb.BookIdRequest(id=1234))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('GetUser failed with {0}: {1}'.format(e.code(), e.details()))

def main():
    """Python Client for Books"""

    # Create channel and stub to server's address and port.
    channel = grpc.insecure_channel('0.0.0.0:50051')
    stub = books_pb2_grpc.BookServiceStub(channel)

    print("-------------- GetBook --------------")
    get_book(stub)
    print("-------------- ListBooks --------------")
    list_books(stub)

if __name__ == '__main__':
    main()
