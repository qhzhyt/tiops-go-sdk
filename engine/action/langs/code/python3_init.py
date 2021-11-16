

def call_func(func):
    def inner(data):
        func(data)

    return inner


def call_batch_func(func):
    def inner(data):
        func(data)
    return inner
