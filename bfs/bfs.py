#!/usr/bin/env python

import sys
import collections
import Queue

def ReadFile(filename):
  try:
    handle = open(filename)
  except IOError:
    print "could not read file."
    sys.exit(-1)
  lines = handle.readlines()
  bfs_text = collections.OrderedDict()
  for line in lines:
    bfs_text[int(line.split()[0])] = [int(i) for i in line.split()[1:]]
  return bfs_text


class BFS(object):

  def __init__(self, graph):
    self.graph = graph
    self.q = Queue.Queue()
    self.done = []
    self.root_node = self.graph.keys()[0]
    print "Root Node %d" % self.root_node
    self.q.put(self.root_node)

  def bfs(self):
    count = 0
    while not self.q.empty():
      node = self.q.get()
      adjacencies = self.graph[node]
      for node in adjacencies:
        if node not in self.done:
          self.done.append(node)
          self.q.put(node)
      count += 1
    print count


def main():
  graph = ReadFile("input.txt")
  b = BFS(graph)
  b.bfs()
  print sorted(b.done)


if __name__ == '__main__':
  main()
