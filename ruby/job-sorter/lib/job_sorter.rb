# frozen_string_literal: true

require 'ostruct'

# Sort jobs with dependency
class JobSorter
  def initialize(string)
    @tree   = Tree.new
    @input  = Input.new(string)

    build_tree_base
    build_tree_dependencies

    raise CircularDependencyException, 'Circular dependency' unless valid?
  end

  def to_s
    tree.present
  end

  private

  attr_reader :tree, :input

  def valid?
    tree.nodes.count == input.nodes.count
  end

  # Add nodes with no dependency.
  def build_tree_base
    input.nodes_with_no_dependencies.each do |node|
      tree.add_node(node.task)
    end
  end

  # Add dependencies.
  def build_tree_dependencies
    loop do
      keep_running = false

      input.nodes_those_depend_of(tree.nodes).each do |item|
        keep_running = tree.add_node_with_dependency(item.to_h)
      end

      break unless keep_running
    end
  end

  # Handle nodes
  class Tree
    attr_reader :nodes

    def initialize
      @nodes = []
    end

    def add_node(node)
      @nodes << node
      # handler = NodeHandler.new(node, nodes)

      # handler.handle
    end

    def add_node_with_dependency(task:, dependency:)
      index = index_of(dependency)

      @nodes = nodes[0..index] + [task] + nodes[index + 1..-1]
    end

    def present
      nodes.join(',')
    end

    private

    def index_of(value)
      nodes.index { |item| item == value }
    end
  end

  # Instance of node
  class Node < OpenStruct
    def initialize(raw)
      key, value = raw.to_a.flatten

      raise SelfDependencyException, 'Self dependency' if (key == value) && !key.nil?

      super(task: key, dependency: value)
    end
  end

  # Handle inputed string
  class Input
    attr_reader :nodes

    def initialize(string)
      @nodes = string.to_s.split(/\n/).each_with_object([]) do |item, response|
        key, value = item.strip.split(/=>/)
        response << Node.new(key => value)
      end
    end

    def nodes_with_no_dependencies
      nodes.select { |node| node.dependency.nil? }
    end

    def nodes_those_depend_of(items)
      nodes.select { |node| items.include?(node.dependency) }
           .reject { |item| items.include?(item.task) }
    end
  end

  # Exception for self dependency
  class SelfDependencyException < StandardError; end

  # Exception for circular dependency
  class CircularDependencyException < StandardError; end
end
