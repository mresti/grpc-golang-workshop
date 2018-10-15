import grpc

import books_pb2 as pb
import books_pb2_grpc

def list_books(stub):
    # Exception handling.
    try:
        # List books
        books = stub.List(pb.Empty())
        print(books)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('ListBooks failed with {0}: {1}'.format(e.code(), e.details()))

def get_book(stub):
    # Exception handling.
    try:
        # Get book
        response = stub.Get(pb.BookIdRequest(id=123))
        print(response)

        response = stub.Get(pb.BookIdRequest(id=1234))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('GetBook failed with {0}: {1}'.format(e.code(), e.details()))

def insert_book(stub):
    # Exception handling.
    try:
        # Insert book
        response = stub.Insert(pb.Book(id=12345, title='My book', author='me'))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('InsertBook failed with {0}: {1}'.format(e.code(), e.details()))

def delete_book(stub):
    # Exception handling.
    try:
        # Insert book
        response = stub.Delete(pb.BookIdRequest(id=12345))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print('DeleteBook failed with {0}: {1}'.format(e.code(), e.details()))

def watch_book(stub):
    response = stub.Watch(pb.Empty())
    try:
        for book in response:
            print(book)
    except grpc._channel._Rendezvous as err:
        print(err)


def GetClient():
    # Create channel and stub to server's address and port.
    channel = grpc.insecure_channel('0.0.0.0:50051')
    stub = books_pb2_grpc.BookServiceStub(channel)
    return stub

def main():
    """Python Client for Books"""
    stub = GetClient()
    print("-------------- GetBook --------------")
    get_book(stub)
    print("-------------- ListBooks --------------")
    list_books(stub)
    print("-------------- InsertBook --------------")
    insert_book(stub)
    print("-------------- ListBooks --------------")
    list_books(stub)
    print("-------------- DeletetBook --------------")
    delete_book(stub)
    print("-------------- ListBooks --------------")
    list_books(stub)
    #print("-------------- WatchBooks --------------")
    #watch_book(stub)

if __name__ == '__main__':
    main()
