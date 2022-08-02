# switch to linkedlists... eventually
class LinkedList:
    def __init__(head=None):
        self.head = head

class Node:
    def __init__(head=False, next=None, prev=None):
        self.head = head
        self.next = next
        self.prev = prev


        
class Interpreter:

    def __init__(self):
        self.pointers = [0]
        self.pointer_index = 0
        self.code = input("test code:")
    
    def move_left(self):
        if self.pointer_index != 0:
            self.pointer_index -= 1
        else:
            print("bad pointer call, waiting for linked list support")
            self.pointer_index = -1

    def move_right(self):
        if self.pointer_index == (len(self.pointers) -1):
            self.pointers.append(0)
            self.pointer_index += 1
        else:
            self.pointer_index += 1

    def add_one(self):
        self.pointers[self.pointer_index] += 1
        
    def sub_one(self):
        self.pointers[self.pointer_index] -= 1

    def run_loop(self,loop_code):
        while True:
            if self.pointers[self.pointer_index] == 1:
                return 
            else:
                self.enterpret(loop_code)

    def show_ascii(self):
        print(chr(self.pointers[self.pointer_index]))

    def get_input(self):
        self.pointers[self.pointer_index] += input()
        
    def enterpret(self, code):
        for index in range(len(code)):
            if code[index] == ">":
                self.move_right()
            elif code[index] == "<":
                self.move_left()
            elif code[index] == "+":
                self.add_one()
            elif code[index] == "-":
                self.sub_one()
            elif code[index] == "]":
                break
            elif code[index] == "[":
                self.run_loop(code[(index+1):])
            elif code[index] == ",":
                self.get_input()
            elif code[index] == ".":
                self.show_ascii()
        print(self.pointers)

    def run(self):
        self.enterpret(self.code)



ii = Interpreter()
ii.run()
print(ii.pointers)
