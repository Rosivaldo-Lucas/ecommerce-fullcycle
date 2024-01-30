import {
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { OrderStatus } from './order-status.enum';
import { OrderItem } from './order-item.entity';

export type CreateOrderCommand = {
  client_id: number;
  items: {
    quantity: number;
    price: number;
    product_id: string;
  }[];
};

@Entity()
export class Order {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ type: 'decimal', precision: 10, scale: 2 })
  total: number;

  @Column()
  client_id: number;

  @Column()
  status: OrderStatus = OrderStatus.PENDING;

  @CreateDateColumn()
  creadted_at: Date;

  @OneToMany(() => OrderItem, (item) => item.order, { cascade: ['insert'] })
  items: OrderItem[];

  static create(input: CreateOrderCommand): Order {
    const order = new Order();

    order.cliente_id = input.client_id;

    order.items = input.items.map((item) => {
      const orderItem = new OrderItem();
      orderItem.product_id = item.product_id;
      orderItem.quantity = item.quantity;
      orderItem.price = item.price;

      return orderItem;
    });

    order.total = order.items.reduce((sum, item) => {
      return sum + item.price * item.quantity;
    }, 0);

    return order;
  }
}
